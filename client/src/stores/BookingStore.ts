import { makeSimpleAutoObservable } from "utils/mobx-extensions"
import { action } from "mobx";
import { api } from "utils/api"
import BasicStore from "./BasicStore"
import { FlightType } from "./FlightStore"
import { TimeLike } from "fs";

type PassengerType = {
    "birthdate": TimeLike,
    "firstname": string,
    "lastname": string,
    "passportCountryId": number,
    "passportNumber": string,
    "phone": string
}

class BookingStore extends BasicStore {
    constructor(...args: any[]) {
        super(...args);
        makeSimpleAutoObservable(this);
    }

    outbound: FlightType | Record<string, never> = {};
    return: FlightType | Record<string, never> = {};
    passengers: PassengerType[] = [];

}

export default BookingStore;