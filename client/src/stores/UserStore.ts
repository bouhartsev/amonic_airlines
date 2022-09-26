import { makeSimpleAutoObservable } from "utils/mobx-extensions"
import { action } from "mobx";
import api, { setAuthToken } from "utils/api"
import BasicStore from "./BasicStore"

type userType = { id: number, name: string, role: string };

class UserStore extends BasicStore {
    constructor(...args: any[]) {
        super(args);
        makeSimpleAutoObservable(this);
    }

    isLogged = false;
    userData = {} as userType;
    users = [] as userType[];

    login = (username: string, password: string) => {
        // temp
        // this.isLogged = true;

        return api.post("/auth/sign-in", { username, password })
            .then((response) => {
                this.isLogged = true;
                localStorage.setItem("jwtToken", response.data.token);
                setAuthToken(response.data.token);
                // get user data
            })
            .catch((err) => {
                this.status = "error";
                switch (err.code) {
                    case "AttemptsExceeded":
                        let tryAfter = 10;
                        let attemptsTimer = setInterval(action(() => {
                            this.error = `You entered the credentials incorrectly 3 times. Next attempt after ${tryAfter--} seconds.`
                            if (tryAfter === 0) {
                                clearInterval(attemptsTimer);
                                if (this.status === "error") {
                                    this.error = "";
                                    this.status = "initial";
                                }
                            }
                        }), 1000);
                        break;
                    case "invalid_credentials":
                    case "user:disabled":
                        this.error = err.message;
                        break;
                    default:
                        throw err;
                }
            });
    }

    logout = () => {
        // temp
        this.isLogged = false;

        // return api.post("/auth/sign-out")
        //     .then((response) => {
        //         this.isLogged = false;
        //         localStorage.removeItem("jwtToken");
        //         setAuthToken();
        //         // remove user data
        //     })
        //     .catch((err) => { this.status = "error"; });
    }
};

export default UserStore;