// import { makeAutoObservable } from "mobx"

abstract class BasicStore {
    rootStore
    status = "initial" as "initial"|"pending"|"success"|"error"|"forbidden";
    error = null as string | null;

    constructor(rootStore?: any) {
        this.rootStore = rootStore;
    }
}

export default BasicStore