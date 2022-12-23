import { Typography, Grid } from "@mui/material";
import Button from "@mui/material/Button";
// import {
//   useForm,
//   Controller,
//   SubmitHandler,
//   useFormState,
// } from "react-hook-form";
import {
  useForm,
  FormContainer,
  TextFieldElement,
  AutocompleteElement,
  DatePickerElement,
} from "react-hook-form-mui";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";

import { observer } from "mobx-react-lite";
import { runInAction } from "mobx";
import { useStore } from "stores";
import { PassengerType } from "stores/BookingStore";
import PassengersList from "./PassengersList";
import { useEffect } from "react";

const BookingConf = () => {
  const { bookingStore, flightStore } = useStore();
  const outboundSchedule = flightStore.scheduleByID(
    bookingStore.outbound?.scheduleId
  );
  const returnSchedule = flightStore.scheduleByID(
    bookingStore.return?.scheduleId
  );
  const handleSubmit = (data: PassengerType) => {
    runInAction(() => bookingStore.passengers.push(data));
  };

  const formContext = useForm<PassengerType>();

  useEffect(() => {
    bookingStore?.getCountries();

    return () => {};
  }, []);

  return (
    <>
      <Typography variant="subtitle1" sx={{my:2}}>Outbound flight details</Typography>
      <Grid
        container
        justifyContent="space-between"
        wrap="wrap"
        rowSpacing={2}
        columnSpacing={2}
      >
        <Grid item>
          From: <b>{outboundSchedule?.from}</b>
        </Grid>
        <Grid item>
          To: <b>{outboundSchedule?.to}</b>
        </Grid>
        <Grid item>
          Cabin type: <b>{bookingStore.outbound?.cabinTypeId}</b>
        </Grid>
        <Grid item>
          Date: <b>{outboundSchedule?.date.toLocaleDateString()}</b>
        </Grid>
        <Grid item>
          Flight number: <b>{outboundSchedule?.flightNumber}</b>
        </Grid>
      </Grid>

      {!!returnSchedule && <>Here is a copy</>}

      <LocalizationProvider dateAdapter={AdapterDateFns}>
        <FormContainer formContext={formContext} onSuccess={handleSubmit}>
          <Typography sx={{ my: 2 }} variant="subtitle1">
            Passenger details
          </Typography>
          <Grid
            container
            rowSpacing={2}
            columnSpacing={2}
            justifyContent="space-between"
            alignItems="start"
          >
            <Grid item xs={6} sm={4}>
              <TextFieldElement
                required
                name="firstName"
                label="First name"
                fullWidth
              />
            </Grid>
            <Grid item xs={6} sm={4}>
              <TextFieldElement
                required
                name="lastName"
                label="Last name"
                fullWidth
              />
            </Grid>
            <Grid item xs={6} sm={4}>
              <DatePickerElement
                required
                name="birthdate"
                label="Birthday"
                inputProps={{ fullWidth: true }}
              />
            </Grid>
            <Grid item xs={6} sm={4}>
              <TextFieldElement
                required
                name="passportNumber"
                label="Passport number"
                fullWidth
              />
            </Grid>
            <Grid item xs={6} sm={4}>
              <AutocompleteElement
                required
                matchId
                name="passportCountryId"
                label="Passport country"
                autocompleteProps={{
                  getOptionLabel: (opt: any) => opt.name,
                }}
                options={bookingStore.countries}
              />
            </Grid>
            <Grid item xs={6} sm={4}>
              <TextFieldElement required name="phone" label="Phone" fullWidth />
            </Grid>
          </Grid>
          <Button type="submit" variant="contained" fullWidth sx={{ my: 2 }}>
            Add passenger
          </Button>
        </FormContainer>
      </LocalizationProvider>

      {!!bookingStore.passengers.length && <PassengersList />}
    </>
  );
};

export default observer(BookingConf);
