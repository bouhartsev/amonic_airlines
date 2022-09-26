// import { makeAutoObservable } from "mobx"

abstract class BasicStore {
    rootStore
    status = "initial"; // initial, pending, success, error
    error = null as string | null;

    constructor(rootStore?: any) {
        this.rootStore = rootStore;
    }
}

export default BasicStore