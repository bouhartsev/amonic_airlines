import { makeSimpleAutoObservable } from "utils/mobx-extensions"
import { action, runInAction } from "mobx";
import { api } from "utils/api"
import BasicStore from "./BasicStore"
import { TimeLike } from "fs";

export type FlightType = {
    "id": number | string,
    "date": Date,
    "time": TimeLike,
    "confirmed": boolean,
    "flightNumber": number,
    "from": string, //string??
    "to": string //string??
    "aircraft": string,
    "economyPrice": number,
    "businessPrice"?: number,
    "firstClassPrice"?: number,
    "emptySeats"?: number,
}

type AirportType = {
    "IATACode": string,
    "countryId": number,
    "id": number,
    "name": string,
}

class FlightStore extends BasicStore {
    constructor(...args: any[]) {
        super(...args);
        makeSimpleAutoObservable(this);
    }

    airports: AirportType[] = [];
    schedules: FlightType[] = [];
    upload = { progress: 0, results: <Record<string, never>>{} };

    scheduleByID = (id: number | string) => this.schedules.find((item) => item.id == id);
    airportByID = (id: number | string) => this.airports.find((item) => item.id == id);

    getSchedules = () => {
        this.status = "pending";
        return api.get("/schedules")
            .then((response: any) => {
                runInAction(() => {
                    this.status = "success";
                    this.error = "";
                    this.schedules = response.data.schedules;
                });
            })
            .catch((err) => { this.status = "error"; throw err; });
    }
    getAirports = () => {
        this.status = "pending";
        return api.get("/airports")
            .then((response: any) => {
                this.status = "success";
                this.error = "";
                this.airports = response.data.airports;
            })
            .catch((err) => { this.status = "error"; throw err; });
    }
    switchConfirm = (flightId: FlightType["id"]) => {
        this.status = "pending";
        return api.post("/users/" + flightId + "/switch-status").then(() => {
            this.status = "success";
            const ind = this.schedules.findIndex((item) => item?.id == flightId);
            this.schedules[ind].confirmed = !this.schedules[ind].confirmed;
        })
    }
    updateSchedule = (data: FlightType) => {
        this.status = "pending";
        const flightId = data.id
        return api.patch("/schedules/" + flightId, (({ date, time, economyPrice }) => ({ date, time, economyPrice }))(data))
            .then((response: any) => {
                this.status = "success";
                this.error = "";
                this.getSchedules();
                return;
            })
            .catch((err) => { this.status = "error"; throw err; });
    }
    uploadSchedules = (files: any) => {
        if (!files || !files.length) return;
        const formData = new FormData();
        formData.append("file", files[0]);
        return api.post("/schedules/upload", formData, {
            headers: {
                'Content-Type': 'multipart/form-data'
            },
            onUploadProgress: (e) => { action(() => this.upload.progress = e.loaded) },
        }).then((response: any) => {
            this.status = "success";
            this.error = "";
            this.upload.results = response.data
        })
            .catch((err) => { this.status = "error"; throw err; });
    }
}

export default FlightStore;