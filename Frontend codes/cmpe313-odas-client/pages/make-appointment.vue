<template>
  <v-container>
    <v-row class="text-center">
      <b>Make An Appointment > &nbsp;&nbsp;</b>
      <i>Search Available Appointments</i>
      <br />
    </v-row>
    <v-row>
      <v-col class="text-start">
        <br />
        |> Search By Date
        <SearchByDate />
        <v-btn color="primary"
          ><v-icon>mdi-clipboard-search</v-icon>&nbsp;&nbsp;SEARCH</v-btn
        >
      </v-col>
      <v-col class="text-start">
        <br />
        |> Search By Hospital
        <SearchByHospital />
        <v-btn color="primary"
          ><v-icon>mdi-clipboard-search</v-icon>&nbsp;&nbsp;SEARCH</v-btn
        >
      </v-col>
      <v-col class="text-start">
        <br />
        |> Search By Doctor
        <SearchByDoctor />
        <v-btn color="primary"
          ><v-icon>mdi-clipboard-search</v-icon>&nbsp;&nbsp;SEARCH</v-btn
        >
      </v-col>
    </v-row>
    <div><br/></div>
    <v-row>
      <v-data-table loading="true" class="elevation-1 xtable"></v-data-table>
      <v-spacer />
      <v-btn x-large color="red" class="ma-10"
        ><v-icon x-large>mdi-book-plus</v-icon>&nbsp;&nbsp;BOOK THE CHOSEN
        APPOINTMENT</v-btn
      >
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'MakeAppointmentPage',
  layout: 'dash',
  middleware({ store, redirect }) {
    if (store.state.auth.user.profile.role !== 'Patient') {
      return redirect('/dashboard')
    }
  },
}
</script>

<style scoped>
.xtable {
  min-width: 70%;
}
</style>
