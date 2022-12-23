import { makeSimpleAutoObservable } from "utils/mobx-extensions"
import { action, runInAction } from "mobx";
import { api, setAuthToken } from "utils/api"
import BasicStore from "./BasicStore"
import jwt_decode from "jwt-decode";

export const roles = ["administrator", "office user"] as const;
export const rolesOptions = roles
  .map((el, ind) => ({ id: (ind + 1).toString(), label: el }))
  .reverse();
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

type ProfileType = {
    "iat": number,
    "exp": number,
    "LastLoginErrorDatetime": string,
    "numberOfCrashes": number,
    "userLogins": Record<string,
        {
            "id": string,
            "error": string,
            "loginTime": string,
            "logoutTime": string,
            "timeSpent": string,
        }
    >[]
}

export type ReportType = {
    "reason": string,
    "softwareCrash": boolean,
    "systemCrash": boolean,
}

class UserStore extends BasicStore {
    constructor(...args: any[]) {
        super(...args);
        makeSimpleAutoObservable(this);

        const currToken = localStorage.getItem("jwtToken");
        if (currToken) this.setAuth(currToken);
    }

    isLogged: boolean = false;
    userData: UserType | Record<string, never> = {};
    profileData: ProfileType | Record<string, never> = {};
    users: UserType[] = [];
    offices: OfficeType[] = [];

    officeByID = (id: number | string) => this.offices.find((item) => item.id == id);
    userByID = (id: number | string) => this.users.find((item) => item?.id == id);

    setAuth = (token: string) => {
        const tokenData: any = jwt_decode(token);
        const { user, ...profile } = tokenData;
        if (profile.exp * 1000 - Date.now() <= 0) return; // check if token is valid
        this.isLogged = true;
        setAuthToken(token);
        this.userData = { ...user, role: roleByID(user.roleId) };
        this.profileData = { ...this.profileData, ...profile };
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

    getProfile = (userId: UserType["id"]) => {
        this.status = "pending";
        return api.get("/users/" + userId + "/logins")
            .then((response: any) => {
                runInAction(() => {
                    this.status = "success";
                    this.error = "";
                    const data = response.data;
                    data.userLogins = data.userLogins.map((el: ProfileType["userLogins"][number], ind: number) => ({ ...el, id: ind }));
                    this.profileData = { ...this.profileData, ...data };
                });
            })
            .catch((err) => { this.status = "error"; throw err; });
    }
    getUsers = () => {
        this.status = "pending";
        return api.get("/users")
            .then((response: any) => {
                runInAction(() => {
                    this.status = "success";
                    this.error = "";
                    this.users = response.data.users;
                });
            })
            .catch((err) => { this.status = "error"; throw err; });
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
    
    report = (data: ReportType) => {
        this.status = "pending";
        return api.post("/auth/report", data).then(() => {
            this.status = "success";
            this.getProfile(this.userData.id);
        })
    }
};

export default UserStore;