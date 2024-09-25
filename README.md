# graphql-db-api-example  

本项目提供了 [graphql-db-api](https://github.com/lujin3/graphql-db-api) 的 go demo 示例, 用于适配 superset 的 graphql db 数据源

## 添加新接口

按照以下步骤在您的 GraphQL 服务器中添加新的 API 端点：

1. **定义 schema**  
   在 `graph/schema.graphql` 中更新您的新 schema 定义。例如：

   ```graphql
   type User {
     id: ID!
     name: String!
   }
   ```

2. **生成 gqlgen 配置和模型**  
    运行以下命令以生成所需的文件：

    ```bash
    go run github.com/99designs/gqlgen generate
    ```  

3. **开发 resolver**  
    在 `graph/schema.resolvers.go` 中实现您新 API 的逻辑。该文件包含负责获取和返回数据的 resolver。

## Quick start

### Start the graphql server  

```bash
go run github.com/99designs/gqlgen generate
```
