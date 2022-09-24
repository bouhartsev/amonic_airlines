import { createContext, useContext } from "react";
import UserStore from "./UserStore";

const store = {
  userStore: new UserStore(),
};

export const StoreContext = createContext(store);

export const useStore = () => {
  return useContext<typeof store>(StoreContext);
};

export default store;