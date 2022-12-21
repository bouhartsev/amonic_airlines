import { makeSimpleAutoObservable } from "utils/mobx-extensions"
import { action, runInAction } from "mobx";
import { api, setAuthToken } from "utils/api"
import BasicStore from "./BasicStore"
import jwt_decode from "jwt-decode";

export const roles = ["administrator", "office user"] as const;
export const roleByID = (roleId: userType["roleId"]) => roles[Number(roleId) - 1];

export type userType = {
    id: number | string,
    role?: typeof roles[number],
    active?: boolean,
    age?: number,
    birthdate: string,
    email: string,
    firstName: string,
    lastName: string,
    officeId: number | string,
    roleId: number | string,
    password?: string,
};

type officeType = { id: number, title: string };

class UserStore extends BasicStore {
    constructor(...args: any[]) {
        super(...args);
        makeSimpleAutoObservable(this);

        const currToken = localStorage.getItem("jwtToken");
        if (currToken) this.setAuth(currToken);
    }

    isLogged: boolean = false;
    userData: userType | Record<string, never> = {};
    users: userType[] = [];
    offices: officeType[] = [];

    officeByID = (officeId: number | string) => this.offices.find((item) => item.id == officeId);
    userByID = (userId: number | string) => this.users.find((item) => item?.id == userId);

    setAuth = (token: string) => {
        const tokenData: any = jwt_decode(token);
        console.log(tokenData);
        if (Date.now() >= tokenData.exp * 1000) return; // check if token is valid
        this.isLogged = true;
        setAuthToken(token);
        this.userData = { ...tokenData.user, role: roleByID(tokenData.user.roleId) };
    }

    login = (username: string, password: string) => {
        this.status = "pending";

        return api.post("/auth/sign-in", { login: username, password })
            .then((response: any) => {
                this.status = "success";
                this.error = "";
                localStorage.setItem("jwtToken", response.data.token);
                this.setAuth(response.data.token);
            })
            .catch((err: any) => {
                this.status = "error";

                switch (err.response?.data?.code) {
                    case "invalid_credentials:series": // not exactly correct on backend
                        console.log(err.response.data)
                        this.status = "forbidden";
                        let nextTry = new Date(err.response.data.details?.NextTry).getTime();
                        let attemptsTimer = setInterval(action(() => {
                            let tryAfter = Math.ceil((nextTry - Date.now()) / 1000);
                            this.error = `You entered the credentials incorrectly 3 times. Next attempt after ${tryAfter} seconds.`
                            if (tryAfter <= 0) {
                                clearInterval(attemptsTimer);
                                if (this.status === "error" || this.status === "forbidden") {
                                    this.error = "";
                                    this.status = "initial";
                                }
                            }
                        }), 1000);
                        break;
                    case "invalid_credentials":
                    case "user:disabled":
                        this.error = err.response.data.message;
                        break;
                    default:
                        throw err;
                }
            });
    }

    logout = () => {
        return api.post("/auth/sign-out")
            .then((response) => {
                this.isLogged = false;
                localStorage.removeItem("jwtToken");
                setAuthToken();
                // remove user data
            })
            .catch((err) => { this.status = "error"; });
    }

    getUsers = () => {
        this.status = "pending";
        return api.get("/users")
            .then((response: any) => {
                runInAction(() => {
                    this.status = "success";
                    this.error = "";
                    this.users = response.data.users;//.map((el: any) => ({ ...el, active: false }));
                });
            })
            .catch((err) => { this.status = "error"; }); // throw err;
    }

    getOffices = () => {
        this.status = "pending";
        return api.get("/offices")
            .then((response: any) => {
                this.status = "success";
                this.error = "";
                this.offices = response.data.offices;
            })
            .catch((err) => { this.status = "error"; throw err; });
    }
    addUser = (data: userType) => {
        this.status = "pending";
        return api.post("/users", (({ email, firstName, lastName, officeId, birthdate, password }) => ({ email, firstName, lastName, officeId, birthdate, password }))(data))
            .then((response: any) => {
                this.status = "success";
                this.error = "";
                this.users.push(response.data.user);
            })
            .catch((err) => {
                this.status = "error";
                if (err.response?.data?.code === "user.email:already_taken") this.error = err.response.data.message;
                else throw err;
            });
    }
    updateUser = (data: userType) => {
        this.status = "pending";
        const userId = data.id
        return api.patch("/users/" + userId, (({ email, firstName, lastName, officeId, roleId }) => ({ email, firstName, lastName, officeId, roleId: Number(roleId) }))(data))
            .then((response: any) => {
                this.status = "success";
                this.error = "";
                const ind = this.users.findIndex((item) => item?.id == userId);
                this.users.splice(ind, 1,response.data.user);
                // console.log(response.data, data.roleId);
                return;
            })
            .catch((err) => { this.status = "error"; throw err; });
    }
};

export default UserStore;