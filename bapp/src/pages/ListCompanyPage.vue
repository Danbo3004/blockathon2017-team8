<template>
    <v-container>
        <h3 class="my-3">Buy Tokens</h3>
        <v-layout row>
    <v-flex xs12 sm6 offset-sm3>
      <v-card>
        <v-toolbar color="indigo" dark>
          <v-toolbar-title>Top Market</v-toolbar-title>
        </v-toolbar>
        <v-list>
          <v-list-tile avatar v-for="(item, index) in items" v-bind:key="item.title" @click="" :to="{name: 'buy_tokens', params: {id: index}}">
            <v-list-tile-action>
              <v-icon v-if="item.icon" color="pink">star</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title v-text="item.title"></v-list-tile-title>
            </v-list-tile-content>
            <v-list-tile-avatar>
              <img v-bind:src="item.avatar"/>
            </v-list-tile-avatar>
          </v-list-tile>
        </v-list>
      </v-card>
    </v-flex>
  </v-layout>
    </v-container>
</template>

<script>
export default {
  data() {
    return {
      items: [
        {
          icon: true,
          title: "Every Coin",
          avatar: "/public/1.png"
        },
        { icon: true, title: "F*nny Coin", avatar: "/public/2.png" },
        { title: "Tiki", avatar: "/public/3.png" },
        { title: "Boo Coffee", avatar: "/public/4.png" },
        { title: "Next Movement", avatar: "/public/5.png" },
        { title: "Leflair", avatar: "/public/6.png" },
      ]
    };
  },

  mounted() {}
};
</script>
