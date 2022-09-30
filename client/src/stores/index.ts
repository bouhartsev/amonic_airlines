import { createContext, useContext } from "react";
import UserStore from "./UserStore";
import FakeStore from "./FakeStore";

class RootStore {
  fakeStore
  userStore

  constructor() {
      this.fakeStore = new FakeStore(this);
      this.userStore = new UserStore(this);
  }
}

const store = new RootStore();

export const StoreContext = createContext(store);

export const useStore = () => {
  return useContext<typeof store>(StoreContext);
};

export default store;