import { makeSimpleAutoObservable } from "utils/mobx-extensions"
import { action, runInAction } from "mobx";
import api, { setAuthToken } from "utils/api"
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
                            let tryAfter = Math.ceil((nextTry - Date.now())/1000);
                            this.error = `You entered the credentials incorrectly 3 times. Next attempt after ${tryAfter} seconds.`
                            if (tryAfter <= 0) {
                                clearInterval(attemptsTimer);
                                this.rootStore.fakeStore.resetLoginAtt();
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
        // temp
        this.isLogged = false;
        localStorage.removeItem("jwtToken");
        setAuthToken();

        // return api.post("/auth/sign-out")
        //     .then((response) => {
        //         this.isLogged = false;
        //         localStorage.removeItem("jwtToken");
        //         setAuthToken();
        //         // remove user data
        //     })
        //     .catch((err) => { this.status = "error"; });
    }

    getUsers = () => {
        this.status = "pending";
        return api.get("/users")
            .then((response: any) => {
                runInAction(() => {
                    this.status = "success";
                    this.users = response.data.users;//.map((el: any) => ({ ...el, active: false }));
                });
            })
            .catch((err) => { this.status = "error"; throw err; });
    }

    getOffices = () => {
        this.status = "pending";
        return api.get("/offices")
            .then((response: any) => {
                this.status = "success";
                this.offices = response.data.offices;
            })
            .catch((err) => { this.status = "error"; throw err; });
    }
};

export default UserStore;