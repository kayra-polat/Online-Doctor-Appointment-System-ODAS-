<template>
  <v-app dark>
    <v-app-bar fixed app>
      <v-btn to="/" class="ml-1 rotatorY" icon>
        <v-icon>mdi-hospital-building</v-icon>
      </v-btn>
      <v-toolbar-title class="rotatorY" v-text="title" />
      <v-spacer />
      <v-btn to="/register" color="error">
        <v-icon>mdi-account-plus</v-icon>
        <span>REGISTER</span>
      </v-btn>
      &nbsp;
      <v-btn to="/login" color="black">
        <v-icon>mdi-login</v-icon>
        <span>SIGN IN</span>
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
    <v-footer app>
      <v-spacer />
      <span>ODAS LLC. &copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: 'DefaultLayout',
  data() {
    return {
      title: 'Online Doctor Appointment System',
    }
  },
}
</script>

<style>
.rotatorY {
  transform: rotateY(560deg);
  animation: turn 2s ease-out forwards 1s;
}

@keyframes turn {
  100% {
    transform: rotateY(0deg);
  }
}
</style>
