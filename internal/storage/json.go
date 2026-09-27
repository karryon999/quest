// storage负责读写数据

//找到用户主目录
//↓
//创建 ~/.quest
//↓
//以后在里面读写 quests.json ，用来保存用户的数据

package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/karryon999/quest/internal/player"
	"github.com/karryon999/quest/internal/quest"
)

// 确保 ~/.quest 存在
func ensureDataDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".quest")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}

//quest

// 得到 ~/.quest/quests.json
func questsFilePath() (string, error) {
	dir, err := ensureDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "quests.json"), nil
}

// Go 数据 → JSON → 文件
// 保存任务
func SaveQuests(quests []quest.Quest) error {
	path, err := questsFilePath()
	if err != nil {
		return err
	}
	//把 Go 里的数据转换成 JSON
	data, err := json.MarshalIndent(quests, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// 文件 → JSON → Go 数据
// 读取任务
func LoadQuests() ([]quest.Quest, error) {
	path, err := questsFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []quest.Quest{}, nil
		}
		return nil, err
	}
	// []quest.Quest是一个quest.Quest类型的切片（切片类似其他语言的数组）
	// 表示一组quest.Quest
	var quests []quest.Quest
	if err := json.Unmarshal(data, &quests); err != nil {
		return nil, err
	}
	return quests, nil
}

//player

// 得到 ~/.quest/player.json
func playerFilePath() (string, error) {
	dir, err := ensureDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "player.json"), nil
}

// 保存玩家
func SavePlayer(p player.Player) error {
	path, err := playerFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// 读取玩家
func LoadPlayer() (player.Player, error) {
	path, err := playerFilePath()
	if err != nil {
		return player.Player{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		//第一次运行的时候player.json不存在这不是error
		if os.IsNotExist(err) {
			return player.Player{}, nil
		}
		return player.Player{}, err
	}
	var p player.Player
	if err := json.Unmarshal(data, &p); err != nil {
		return player.Player{}, err
	}
	return p, nil
}
