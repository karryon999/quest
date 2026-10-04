package storage

import (
	//database/sql 是 Go 标准库提供的数据库统一接口。
	//它不只支持 MySQL，也可以配合 PostgreSQL、SQLite 等数据库使用。
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/karryon999/quest/internal/player"
	"github.com/karryon999/quest/internal/quest"
)

//读取 DSN
//   ↓
//检查有没有配置
//   ↓
//创建 *sql.DB
//   ↓
//Ping MySQL
//   ↓
//成功
//   ↓
//把 db 返回给调用者

//*sql.DB 数据库连接池的管理对象

//连接数据库

func OpenMySQL() (*sql.DB, error) {
	//读取系统环境变量，我们会在终端设置：
	//$env:QUEST_DB_DSN="root:密码@tcp(127.0.0.1:3306)/quest?parseTime=true&loc=Local"
	dsn := os.Getenv("QUEST_DB_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("QUEST_DB_DSN is not set")
	}
	//"mysql" 表示使用刚才注册的 MySQL driver
	//sql.Open() 并不会立刻真正连接 MySQL，而是主要在创建*sql.DB
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//Ping() 才是真正问数据库：
	//“MySQL，你在吗？这个用户名密码能登录吗？这个连接可用吗？”
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// 增

func AddQuestMySQL(db *sql.DB, q quest.Quest) (int64, error) {
	result, err := db.Exec(`INSERT INTO quests (title, xp, boss, completed, created_at, completed_at) VALUES (?, ?, ?, ?, ?, ?)`,
		q.Title,
		q.XP,
		q.Boss,
		q.Completed,
		q.CreatedAt,
		q.CompletedAt,
	)
	if err != nil {
		return 0, err
	}
	//获取刚刚这次 INSERT 操作由数据库自动生成的主键 ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

//cmd命令和sql操作的对应关系：
//add     → INSERT
//list    → SELECT
//delete  → DELETE

//MySQL 返回多行数据
//    ↓
//rows.Next() 一行一行往下走
//    ↓
//rows.Scan() 把当前这一行装进一个 Quest
//    ↓
//append 到 slice
//    ↓
//继续下一行
//    ↓
//全部读完
//    ↓
//return []quest.Quest

// 查所有的数据

func LoadQuestMySQL(db *sql.DB) ([]quest.Quest, error) {
	//db.Query() 返回的是： *sql.Rows
	//你可以把 rows 理解成：
	//数据库返回给你的“一批查询结果 + 一个读取这些结果的游标”。
	rows, err := db.Query(`SELECT id, title, xp, boss, completed, created_at, completed_at FROM quests ORDER BY id`)
	if err != nil {
		return nil, err
	}

	//defer rows.Close() 的意思是：
	//等当前函数快要结束时，再关闭这个查询结果集 rows，释放它占用的数据库资源。
	defer rows.Close()
	var quests []quest.Quest
	for rows.Next() {
		var q quest.Quest
		var completedAt sql.NullTime
		err := rows.Scan(
			&q.ID,
			&q.Title,
			&q.XP,
			&q.Boss,
			&q.Completed,
			&q.CreatedAt,
			&completedAt)
		if err != nil {
			return nil, err
		}
		if completedAt.Valid {
			q.CompletedAt = &completedAt.Time
		}
		quests = append(quests, q)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return quests, nil
}

func LoadPlayerMySQL(db *sql.DB) (player.Player, error) {
	var p player.Player
	var lastCompleteDate sql.NullTime

	err := db.QueryRow(`SELECT xp, current_streak, best_streak, last_completed_date FROM players WHERE id = ?`, 1).Scan(
		&p.XP, &p.CurrentStreak, &p.BestStreak, &lastCompleteDate)
	if err != nil {
		return player.Player{}, err
	}
	if lastCompleteDate.Valid {
		p.LastCompletedDate = &lastCompleteDate.Time
	}
	return p, nil
}

//事务版本的函数

func LoadPlayerTx(tx *sql.Tx) (player.Player, error) {
	var p player.Player
	var lastCompleteDate sql.NullTime
	err := tx.QueryRow(`SELECT xp, current_streak, best_streak, last_completed_date
		FROM players
		WHERE id = ?
		FOR UPDATE`, 1).Scan(
		&p.XP, &p.CurrentStreak, &p.BestStreak, &lastCompleteDate)
	if err != nil {
		return player.Player{}, err
	}
	if lastCompleteDate.Valid {
		p.LastCompletedDate = &lastCompleteDate.Time
	}
	return p, nil
}

func UpdatePlayerTx(tx *sql.Tx, p player.Player) error {
	_, err := tx.Exec(`
		UPDATE players
		SET
			xp = ?,
			current_streak = ?,
			best_streak = ?,
			last_completed_date = ?
		WHERE id = ?`,
		p.XP,
		p.CurrentStreak,
		p.BestStreak,
		p.LastCompletedDate,
		1)
	return err
}

//查询单条数据

func GetQuestByIDMySQL(db *sql.DB, id int) (quest.Quest, bool, error) {
	var q quest.Quest
	var completedAt sql.NullTime
	//db.QueryRow()查询一行
	//db.Query()查询多行
	err := db.QueryRow(`SELECT id, title, xp, boss, completed, created_at, completed_at FROM quests WHERE id=?`, id).Scan(
		&q.ID, &q.Title, &q.XP, &q.Boss, &q.Completed, &q.CreatedAt, &completedAt)
	if err != nil {
		//查询成功执行了，但是没找到任何数据。
		if err == sql.ErrNoRows {
			return quest.Quest{}, false, nil
		}
		return quest.Quest{}, false, err
	}

	if completedAt.Valid {
		q.CompletedAt = &completedAt.Time
	}
	return q, true, nil
}

//事务版本的函数
//在这个事务处理这条 Quest 期间，把这一行锁住，避免另一个操作同时把同一个 Quest 完成两次

func GetQuestByIDTx(tx *sql.Tx, id int) (quest.Quest, bool, error) {
	var q quest.Quest
	var completedAt sql.NullTime
	err := tx.QueryRow(`SELECT id, title, xp, boss, completed, created_at, completed_at
		FROM quests
		WHERE id = ?
		FOR UPDATE`, id).Scan(
		&q.ID,
		&q.Title,
		&q.XP,
		&q.Boss,
		&q.Completed,
		&q.CreatedAt,
		&completedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return quest.Quest{}, false, nil
		}
		return quest.Quest{}, false, err
	}
	if completedAt.Valid {
		q.CompletedAt = &completedAt.Time
	}
	return q, true, nil
}

// 真正把任务标记为完成

func CompleteQuestMySQL(db *sql.DB, id int, completedAt time.Time) error {
	_, err := db.Exec(`UPDATE quests SET completed=?, completed_at=? WHERE id=?`, true, completedAt, id)
	return err
}

func CompleteQuestTx(tx *sql.Tx, id int, completedAt time.Time) error {
	_, err := tx.Exec(`UPDATE quests
		SET completed = ?, completed_at = ?
		WHERE id = ?`, true, completedAt, id)
	return err
}

//删

func DeleteQuestMySQL(db *sql.DB, id int) (bool, error) {
	result, err := db.Exec(`DELETE FROM quests WHERE id = ?`, id)
	if err != nil {
		return false, err
	}

	//获取刚才那条 SQL 实际影响了多少行数据。
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}
