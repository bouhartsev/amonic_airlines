import { makeAutoObservable } from "mobx"
import api, { setAuthToken } from "utils/api"

type userType = { id: number, name: string, role: string };

class UserStore {
    isLogged = false;
    userData = {} as userType;
    users = [] as userType[];
    status = "initial";

    constructor() {
        makeAutoObservable(this);
    }

    login = (login: string, password: string) => {
        return api.post("/auth/sign-in", { login, password })
            .then((response) => {
                this.isLogged = true;
                localStorage.setItem("jwtToken", response.data.token);
                setAuthToken(response.data.token);

            })
            .catch((err) => { this.status = "error"; })
    }

    logout = () => {
        return;
    }
};

export default UserStore;