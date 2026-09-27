# ⚔️ Quest

Quest 是一个使用 Go 编写的 RPG 风格命令行 TODO 应用。

把普通待办事项变成 Quest，完成任务获得 XP、提升等级，并记录连续完成任务的天数。

## 功能

目前已实现：

- 添加任务
- 查看任务
- 完成任务
- 删除任务
- 自定义 XP
- Boss 任务
- XP / Level 系统
- Current Streak / Best Streak
- 玩家统计信息
- JSON 本地持久化

## 使用方式

开发阶段可以直接使用：

```bash
go run . <command>
```

添加任务：

```bash
go run . add "Read a book"
```

指定 XP：

```bash
go run . add "Write report" --xp 50
```

添加 Boss 任务：

```bash
go run . add "Finish project" --xp 100 --boss
```

查看任务：

```bash
go run . list
```

完成任务：

```bash
go run . done 1
```

删除任务：

```bash
go run . delete 1
```

查看玩家属性：

```bash
go run . stats
```

## 示例

```text
⚔️ Today's Quests
1 [x] Write Report +30 XP
2 [ ] Read Book +20 XP 👹
```

玩家属性：

```text
📊 Adventurer Stats
Level: 1
XP: 40
Completed Quests: 3
Current Streak: 1
Best Streak: 1
```

## 等级规则

目前每 100 XP 提升一级：

```text
0 - 99 XP     -> Level 1
100 - 199 XP  -> Level 2
200 - 299 XP  -> Level 3
```

## 数据存储

数据保存在：

```text
~/.quest/
├── quests.json
└── player.json
```

## 项目结构

```text
quest/
├── cmd/
├── internal/
│   ├── quest/
│   ├── player/
│   └── storage/
├── main.go
├── go.mod
└── README.md
```

## 开发检查

```bash
gofmt -w .
go test ./...
go vet ./...
```

## 当前进度

- [x] Core CLI commands
- [x] JSON storage
- [x] XP / Level
- [x] Streak
- [x] Stats
- [ ] Unit tests
- [ ] More RPG features

## 项目地址

https://github.com/karryon999/quest