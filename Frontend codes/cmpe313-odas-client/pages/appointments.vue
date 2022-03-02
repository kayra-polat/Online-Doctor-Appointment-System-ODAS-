<template>
  <v-row>
    <v-snackbar
      v-model="snackbar"
      color="red"
      centered
      transition="slide-x-transition"
    >
      {{ snackbarText }}
      <template #action="{ attrs }">
        <v-btn text v-bind="attrs" @click="snackbar = false"> Close </v-btn>
      </template>
    </v-snackbar>
    <v-col>
      <br />
      <div class="text-center"><b>APPOINTMENTS</b></div>
      <br />
      <v-data-table
        :items-per-page="perPage"
        :headers="keys"
        :items="data"
        :loading="loading"
        loader-height="6"
        item-key="title"
        class="elevation-1"
        @click:row="onRowSelect"
      ></v-data-table>
      <v-overlay v-if="selected != null" :value="overlay" opacity="0.85">
        <v-btn color="red" icon @click="closeOverlay"
          ><v-icon>mdi-close-box</v-icon></v-btn
        >
        <br />
        <v-card>
          <v-card-title
            >EDIT / DELETE THE APPOINTMENT | <v-icon>mdi-pen</v-icon> |
            <v-icon>mdi-delete</v-icon></v-card-title
          >
          <h3>Appointment ID: {{ selected.appointmentID }}</h3>
          <h3>Title: {{ selected.title }}</h3>
          <h3>Patient ID: {{ selected.patientID }}</h3>
          <h3>Doctor ID: {{ selected.doctorID }}</h3>
          <h3>Date: {{ selected.date }}</h3>
          <h3>Hospital: {{ selected.hospital }}</h3>
          <h3>Room: {{ selected.room }}</h3>
        </v-card>
        <br />
        <v-btn color="red">update</v-btn>
      </v-overlay>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'AppointmentsPage',
  layout: 'dash',
  data() {
    return {
      snackbar: false,
      snackbarText: "Couldn't fetch appointments...",
      loading: false,
      perPage: 5,
      keys: [
        {
          text: 'Reason',
          align: 'start',
          sortable: true,
          value: 'title',
        },
        { text: 'Date', value: 'date' },
        { text: 'Hospital', value: 'hospital' },
        { text: 'Room', value: 'room' },
      ],
      data: [],
      overlay: false,
      selected: null,
    }
  },
  async mounted() {
    try {
      await this.fetchAppointments()
    } catch (err) {
      this.snackbar = true
    }
  },
  methods: {
    async fetchAppointments() {
      this.loading = true
      const { data } = await this.$axios.$get('appointments/', {
        headers: {
          Authorization: 'Bearer ' + this.$store.state.auth.user.tokenString,
        },
      })
      this.data = data
      this.loading = false
    },
    onRowSelect(data, e) {
      this.selected = data
      this.overlay = true
    },
    closeOverlay() {
      this.overlay = false
      this.selected = null
    },
  },
}
</script>
