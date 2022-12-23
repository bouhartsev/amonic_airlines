import { makeSimpleAutoObservable } from "utils/mobx-extensions"

class FakeStore {
    constructor(rootStore?: any) {
        this.rootStore = rootStore;
        makeSimpleAutoObservable(this);
    }

    rootStore
    loginAttemts = 0;

    newLoginAtt = () => {
        if (++this.loginAttemts>3) return Promise.reject({code:"invalid_credentials:series"});
        return Promise.resolve();
    }
    resetLoginAtt = () => {
        this.loginAttemts = 0;
    }

}

export default FakeStore;