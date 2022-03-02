<template>
  <v-col>
    <v-row
      justify="center"
      align="center"
      dir="row"
      align-content="center"
      class="mt-12"
    >
      <v-col>
        <v-row justify="center" align="center" dir="row" align-content="center">
          <v-btn
            to="/appointments"
            class="ma-2"
            x-large
            color="grey--darken-4 amber--text text--lighten-1"
          >
            My Appointments&nbsp;&nbsp;
            <v-icon x-large>mdi-clipboard-text-clock</v-icon>
          </v-btn>
        </v-row>
      </v-col>
      <v-divider vertical></v-divider>
      <v-col>
        <v-row justify="center" align="center" dir="row" align-content="center">
          <v-btn
            :to="button2URL"
            class="ma-2"
            x-large
            color="grey--darken-4 amber--text text--lighten-1"
          >
            {{button2Label}}&nbsp;&nbsp;
            <v-icon x-large>mdi-bell-plus</v-icon>
          </v-btn>
        </v-row>
      </v-col>
    </v-row>
    <v-spacer></v-spacer>
    <v-row justify="center" align="center" dir="row" align-content="center" class="mt-6">
      <v-img src="images/patient-dashboard.jpg" aspect-ratio="1.7" contain max-height="60vh"> </v-img>
    </v-row>
  </v-col>
</template>

<script>
export default {
  name: 'DashboardPage',
  layout: 'dash',
  middleware({ store, redirect }) {
    if (!store.state.auth.user) {
      return redirect('/')
    }
  },
  data () {
    return {
      button2Label: (this.$store.state.auth.user.profile.role === 'Patient') 
        ? 'Make An Appointment'
        : 'Arrange Available Time',

      button2URL: (this.$store.state.auth.user.profile.role === 'Patient') 
        ? '/make-appointment'
        : '/arrange-availables'
    }
  }
}
</script>
