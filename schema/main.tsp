import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;

@service(#{
  title: "canalia"
})
@server("https://localhost:8080", "Development server")
namespace CanaliaAPI;

@route("/")
interface Routes {
  @get
  @route("/user")
  @summary("ユーザー情報を取得する")
  @tag("User")
  getUser(): User;
}

model User {
  @doc("ユーザー名")
  name: string;
  
  @doc("メールアドレス")
  email: string;
}