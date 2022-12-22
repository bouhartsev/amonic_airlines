import { createContext, useContext } from "react";
import FakeStore from "./FakeStore";
import UserStore from "./UserStore";
import BookingStore from "./BookingStore"

class RootStore {
  fakeStore
  userStore
  bookingStore

  constructor() {
      this.fakeStore = new FakeStore(this);
      this.userStore = new UserStore(this);
      this.bookingStore = new BookingStore(this);
  }
}

const store = new RootStore();

export const StoreContext = createContext(store);

export const useStore = () => {
  return useContext<typeof store>(StoreContext);
};

export default store;