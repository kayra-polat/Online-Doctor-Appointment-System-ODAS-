<template>
  <v-carousel cycle progress interval="4000" progress-color="red">
    <v-carousel-item>
      <v-row class="fill-height" align="center" justify="center">
        <div class="text-h2">🏥 <b>WELCOME</b> 👋</div>
      </v-row>
    </v-carousel-item>
    <v-carousel-item>
      <v-row class="fill-height" align="center" justify="center">
        <div class="text-h1">📱 We're bringing the world together 🌎</div>
        <div class="text-h4">
          For humanity... For the sake of health rights... 🩺
        </div>
      </v-row>
    </v-carousel-item>
        <v-carousel-item>
      <v-row class="fill-height" align="center" justify="center">
        <div class="text-h2">👩🏿‍⚕️ Your doctor is just one step away...</div>
      </v-row>
    </v-carousel-item>
  </v-carousel>
</template>

<script>
export default {
  name: 'IndexPage',
  middleware({ store, redirect }) {
    if (store.state.auth.user) {
      return redirect('/dashboard')
    }
  },
}
</script>
