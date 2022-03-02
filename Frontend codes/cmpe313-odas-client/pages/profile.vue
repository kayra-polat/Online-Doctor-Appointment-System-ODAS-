<template>
  <v-row
    justify="center"
    align="center"
    dir="row"
    align-content="center"
    class="text-center"
  >
    <v-card min-width="600px">
      <v-list two-line>
        <template v-for="(item, index) in profileItems">
          <v-list-item
            :key="item"
            v-text="index.toUpperCase() + ': ' + item"
          ></v-list-item>
          <span :key="item">
            <v-icon v-if="index === 'userID'">mdi-cancel</v-icon>
            <v-icon v-else>mdi-pen</v-icon>
          </span>
          <v-divider :key="index" :inset="item"></v-divider>
        </template>
      </v-list>
      <v-btn class="ma-4" color="red"> UPDATE PROFILE </v-btn>
    </v-card>
  </v-row>
</template>

<script>
export default {
  name: 'ProfilePage',
  layout: 'dash',
  computed: {
    profileItems() {
      const profile = this.$store.state.auth.user.profile
      delete profile.privilegeLevel

      if (profile.role !== 'Doctor') {
        delete profile.title
        delete profile.expertise
        delete profile.hospital
        delete profile.room
      }
      return profile
    },
  },
}
</script>
