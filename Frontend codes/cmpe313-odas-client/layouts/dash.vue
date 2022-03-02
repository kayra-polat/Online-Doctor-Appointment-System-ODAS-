<template>
  <v-app dark>
    <v-app-bar fixed app>
      <v-btn to="/dashboard" class="ml-1 rotatorY" icon>
        <v-icon>mdi-home</v-icon>
      </v-btn>
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <v-btn to="/profile" color="black">
        <v-icon>mdi-account</v-icon>&nbsp;
        <span>Profile</span>
      </v-btn>
      &nbsp;&nbsp;
      <v-btn to="/faq" color="black">
        <v-icon>mdi-help-box</v-icon>&nbsp;
        <span>FAQ</span>
      </v-btn>
      &nbsp;&nbsp;
      <v-btn to="/contact" color="black">
        <v-icon>mdi-email</v-icon>&nbsp;
        <span>Contact Us</span>
      </v-btn>
      &nbsp;&nbsp;
      <v-btn color="red" @click="logout">
        <v-icon>mdi-logout</v-icon>&nbsp;
        <span>Sign Out</span>
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
  name: 'DashboardLayout',
  data () {
    return {
      title: (this.$store.state.auth.user.profile.role === 'Doctor') 
        ? 'Welcome, Doctor ' + this.$store.state.auth.user.profile.name + ' !'
        : 'Welcome, ' + this.$store.state.auth.user.profile.name + ' !'
    }
  },
  methods: {
    async logout() {
      await this.$store.commit('auth/logout')
      await this.$router.push({ path: '/' })
    }
  }
}
</script>

<style>
.rotatorY {
  transform: rotateY(560deg);
  animation: turn 3s ease-out forwards 1s;
}

@keyframes turn {
  100% {
    transform: rotateY(0deg);
  }
}
</style>
