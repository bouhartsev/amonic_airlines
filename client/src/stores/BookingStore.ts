import { makeSimpleAutoObservable } from "utils/mobx-extensions"
import { action } from "mobx";
import { api } from "utils/api"
import BasicStore from "./BasicStore"
import { FlightType } from "./FlightStore"
import { TimeLike } from "fs";

export type OneWay = {
    "cabinTypeId": number,
    "scheduleId": number,
}

export type PassengerType = {
    "birthdate": Date,
    "firstName": string,
    "lastName": string,
    "passportCountryId": number,
    "passportNumber": string,
    "phone": string
}

type Data = { id: number, name: string }

class BookingStore extends BasicStore {
    constructor(...args: any[]) {
        super(...args);
        makeSimpleAutoObservable(this);
    }

    outbound: OneWay | Record<string, never> = {};
    return: OneWay | Record<string, never> = {};
    passengers: PassengerType[] = [];

    cabinTypes: Data[] = [];
    countries: Data[] = [];

    countryByID = (id: number | string) => this.countries.find(item => item.id == id);
    cabinTypeByID = (id: number | string) => this.cabinTypes.find(item => item.id == id);

    getCabinTypes = () => {
        this.status = "pending";
        return api.get("/cabin-types")
            .then((response: any) => {
                this.status = "success";
                this.error = "";
                this.cabinTypes = response.data.cabinTypes;
            })
            .catch((err) => { this.status = "error"; throw err; });
    }
    getCountries = () => {
        this.status = "pending";
        return api.get("/countries")
            .then((response: any) => {
                this.status = "success";
                this.error = "";
                this.countries = response.data.countries;
            })
            .catch((err) => { this.status = "error"; throw err; });
    }
}

export default BookingStore;