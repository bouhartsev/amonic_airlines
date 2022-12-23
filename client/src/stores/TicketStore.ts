import { makeSimpleAutoObservable } from "utils/mobx-extensions"
import { action } from "mobx";
import { api } from "utils/api"
import { OneWay, PassengerType } from "./BookingStore"
import BasicStore from "./BasicStore"

export type TicketType = {
    "outbound": OneWay,
    "passenger": PassengerType,
    "return": OneWay
}

class TicketStore extends BasicStore {
    constructor(...args: any[]) {
        super(...args);
        makeSimpleAutoObservable(this);
    }
}

export default TicketStore;