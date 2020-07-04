import {
  VuexModule,
  Module,
  Action,
  Mutation,
  getModule
} from "vuex-module-decorators";
import { login, logout, getUserInfo } from "@/api/users";
import { getToken, setToken, removeToken } from "@/utils/cookies";
import store from "@/store";

export interface IUserState {
  token: string;
  username: string;
  avatar: string;
  roles: string[];
}

@Module({ dynamic: true, store, name: "user" })
class User extends VuexModule implements IUserState {
  public token = getToken() || "";
  public username = "";
  public avatar = "";
  public roles: string[] = [];

  @Mutation
  private SET_TOKEN(token: string) {
    this.token = token;
  }

  @Mutation
  private SET_NAME(username: string) {
    this.username = username;
  }

  @Mutation
  private SET_AVATAR(avatar: string) {
    this.avatar = avatar;
  }
  @Mutation
  private SET_ROLES(roles: string[]) {
    this.roles = roles;
  }

  @Action
  public async Login(userInfo: { username: string; password: string }) {
    let { username, password } = userInfo;
    username = username.trim();
    const { data } = await login({ username, password });
    setToken(data.token);
    this.SET_TOKEN(data.token);
  }

  @Action
  public ResetToken() {
    removeToken();
    this.SET_TOKEN("");
    this.SET_ROLES([]);
  }

  @Action
  public async GetUserInfo() {
    if (this.token === "") {
      throw Error("GetUserInfo: token is undefined!");
    }
    const { data } = await getUserInfo({
      /* Your params here */
    });
    if (!data) {
      throw Error("Verification failed, please Login again.");
    }
    const { username, avatar } = data;
    this.SET_ROLES(["admin"]);
    this.SET_NAME(username);
    this.SET_AVATAR(avatar);
  }

  @Action
  public async LogOut() {
    if (this.token === "") {
      throw Error("LogOut: token is undefined!");
    }
    await logout();
    removeToken();
    this.SET_TOKEN("");
    this.SET_ROLES([]);
  }
}

export const UserModule = getModule(User);
