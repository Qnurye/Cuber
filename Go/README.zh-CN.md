# 魔方解算器

## 简介

本程序旨在通过串口通信向 Arduino 微控制器发送指令。它利用一个魔方公式解析器，将魔方旋转序列转换为相应的串口命令。

下位机为双臂魔方机器人。

## 目录结构

项目目录结构组织如下：

```plaintext
├── build
├── config
│   ├── cmd.json
│   ├── delay.json
│   └── formula.json
├── go.mod
├── go.sum
├── main.go
├── pkg
│   ├── config
│   │   ├── cmd.go
│   │   ├── delay.go
│   │   └── formula.go
│   └── rubiksCube
│       ├── parser.go
│       └── README.md
├── README.md
└── test
    ├── cmd_test.go
    ├── delay_test.go
    ├── formula_test.go
    └── parser_test.go
```

- `build`：包含构建产物。
- `config`：存储配置文件。
- `go.mod` 和 `go.sum`：Go 模块文件。
- `main.go`：程序的主入口点。
- `pkg`：包含可重用代码的包目录。
  - `config`：与配置相关的功能。
  - `rubiksCube`：魔方公式解析器和相关文档。
- `test`：项目的单元测试。

## 入门

1. 在系统上安装 Go。
2. 克隆存储库：

   ```bash
   git clone https://github.com/Qnurye/Cuber.git
   ```

3. 进入项目目录：

   ```bash
   cd Cuber/Go
   ```

4. 安装依赖项：

   ```bash
   go mod download
   ```

5. 构建项目：

   ```bash
   go build
   ```

6. 运行可执行文件：

   ```bash
   ./main
   ```

## 配置

`config` 目录包含三个 JSON 文件：

- `cmd.json`：命令配置。
- `delay.json`：延迟配置。
- `formula.json`：公式配置。

编辑这些文件以自定义魔方解算器的行为。

## 使用

程序从 `formula.json` 文件读取魔方公式，并通过串口通信将相应的命令发送到 Arduino。

## 测试

项目的单元测试位于 `test` 目录中。使用以下命令运行测试：

```bash
go test ./...
```

## 贡献

如果您想为该项目做出贡献，请遵循标准的 GitHub 分叉和拉取请求工作流程。

## 许可证

该项目根据 [MIT 许可证](LICENSE) 许可。
