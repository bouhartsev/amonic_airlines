import { makeSimpleAutoObservable } from "utils/mobx-extensions"
import { action } from "mobx";
import { api } from "utils/api"
import BasicStore from "./BasicStore"
import { TimeLike } from "fs";

export type FlightType = {
    "aircraft": string,
    
    "confirmed": boolean,
    "date": Date,
    "economyPrice": number,
    "businessPrice"?: number,
    "firstClassPrice"?: number,
    "flightNumber": number,
    "from": string, //string??
    "id": number,
    "time": TimeLike,
    "to": string //string??
}

class FlightStore extends BasicStore {
    constructor(...args: any[]) {
        super(...args);
        makeSimpleAutoObservable(this);
    }
}

export default FlightStore;