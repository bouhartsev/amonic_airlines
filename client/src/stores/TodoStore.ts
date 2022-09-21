import { makeAutoObservable } from "mobx"

const todoStore = () => {
    return makeAutoObservable({
        list: [] as { title: string; id: number} [],
    });
};

export default todoStore;