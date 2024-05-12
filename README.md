# Go Web 服务器示例

本项目演示了一个简单的 Go HTTP 服务器，包括处理表单提交和提供静态文件服务。

## 功能特性

- 从指定目录提供静态文件服务。
- 处理表单提交并以提交的数据响应。
- 处理特定的路由，对于`/hello` 和 `/form`路径提供特定功能。

## 安装指南

1. 克隆仓库到本地：
    ```bash
    git clone https://your-repository-url.git
    ```
2. 进入项目目录：
   ```bash
   cd your-project-directory
   ```
3. 启动服务器：

   ```bash
   go run main.go

   ```

### 使用说明

- 访问`http://localhost:8080`可以浏览`index.html`静态文件。
- 访问`http://localhost:8080/form.html` 发送POST请求以测试表单提交。
- 访问`http://localhost:8080/hello` 以接收到`"hello!"`的响应。
