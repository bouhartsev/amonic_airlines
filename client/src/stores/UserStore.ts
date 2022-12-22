import { makeSimpleAutoObservable } from "utils/mobx-extensions"
import { action, runInAction } from "mobx";
import { api, setAuthToken } from "utils/api"
import BasicStore from "./BasicStore"
import jwt_decode from "jwt-decode";

export const roles = ["administrator", "office user"] as const;
export const roleByID = (roleId: UserType["roleId"]) => roles[Number(roleId) - 1];

export type UserType = {
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

type OfficeType = { id: number, title: string };

class UserStore extends BasicStore {
    constructor(...args: any[]) {
        super(...args);
        makeSimpleAutoObservable(this);

        const currToken = localStorage.getItem("jwtToken");
        if (currToken) this.setAuth(currToken);
    }

    isLogged: boolean = false;
    userData: UserType | Record<string, never> = {};
    users: UserType[] = [];
    offices: OfficeType[] = [];

    officeByID = (officeId: number | string) => this.offices.find((item) => item.id == officeId);
    userByID = (userId: number | string) => this.users.find((item) => item?.id == userId);

    setAuth = (token: string) => {
        const tokenData: any = jwt_decode(token);
        if (Date.now() >= tokenData.exp * 1000) return; // check if token is valid
        this.isLogged = true;
        setAuthToken(token);
        this.userData = { ...tokenData.user, role: roleByID(tokenData.user.roleId) };
    }
    removeAuth = () => {
        this.isLogged = false;
        localStorage.removeItem("jwtToken");
        setAuthToken();
        // remove all data from all stores
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
        this.status = "pending";
        return api.post("/auth/sign-out").finally(() => {
            this.status = "initial";
            this.removeAuth();
        });
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

    addUser = (data: UserType) => {
        this.status = "pending";
        console.log((new Date(data.birthdate)).toISOString().split('T')[0])
        return api.post("/users", (({ email, firstName, lastName, officeId, birthdate, password }) => ({ email, firstName, lastName, officeId, birthdate: (new Date(birthdate)).toISOString().split('T')[0], password }))(data))
            .then((response: any) => {
                this.status = "success";
                this.error = "";
                console.log(response.data);
                this.users.push(response.data.user);
            })
            .catch((err) => {
                this.status = "error";
                if (err.response?.data?.code === "user.email:already_taken") this.error = err.response.data.message;
                else throw err;
            });
    }

    updateUser = (data: UserType) => {
        this.status = "pending";
        const userId = data.id
        return api.patch("/users/" + userId, (({ email, firstName, lastName, officeId, roleId }) => ({ email, firstName, lastName, officeId, roleId: Number(roleId) }))(data))
            .then((response: any) => {
                this.status = "success";
                this.error = "";
                const ind = this.users.findIndex((item) => item?.id == userId);
                this.users.splice(ind, 1, response.data.user);
                return;
            })
            .catch((err) => { this.status = "error"; throw err; });
    }

    switchActive = (userId: UserType["id"]) => {
        this.status = "pending";
        return api.post("/users/" + userId + "/switch-status").then(() => {
            this.status = "success";
            const ind = this.users.findIndex((item) => item?.id == userId);
            this.users[ind].active = !this.users[ind].active;
        })
    }
};

export default UserStore;