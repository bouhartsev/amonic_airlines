import { makeSimpleAutoObservable } from "utils/mobx-extensions"
import { action } from "mobx";
import api, { setAuthToken } from "utils/api"
import BasicStore from "./BasicStore"

type userType = { id: number, name: string, role: "User"|"Administrator" };

class UserStore extends BasicStore {
    constructor(...args: any[]) {
        super(...args);
        makeSimpleAutoObservable(this);
        
        const currentToken = localStorage.getItem("jwtToken");
        
        if (currentToken) {
            // check if token is Valid
            this.isLogged = true;
            setAuthToken(currentToken);
        }
    }

    isLogged = false;
    userData = {role: "Administrator"} as userType;
    users = [] as userType[];

    login = (username: string, password: string) => {
        this.status = "pending";

        return api.post("/auth/sign-in", { login: username, password })
            .then((response: any) => {
                this.status = "success";
                localStorage.setItem("jwtToken", response.data.token);
                this.isLogged = true;
                setAuthToken(response.data.token);
                // get user data
            })
            .catch((err: any) => {
                this.status = "error";
                console.log(err);
                switch (err.response?.data?.code) {
                    case "invalid_credentials:series":
                        let tryAfter = 10;
                        this.status = "forbidden";
                        let attemptsTimer = setInterval(action(() => {
                            this.error = `You entered the credentials incorrectly 3 times. Next attempt after ${tryAfter--} seconds.`
                            if (tryAfter === 0) {
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
            this.status = "success";
            console.log(response.data);
        })
    }
};

export default UserStore;