<template>
    <v-container>
        <h3 class="my-3">Buy Tokens</h3>
        <v-card color="blue-grey darken-2" class="white--text" style="background: rgb(30, 116, 177)">
            <v-card-title primary-title>
              <div class="headline">{{item.title}}</div>
              <img :src="item.avatar" style="height: 50px; position: absolute; top: 20px; right: 20px"/>
              <div><br>You are about to buy token from <b style="word-break: keep-all">{{item.title}}</b>. How much do you want to buy?</div>
            </v-card-title>
          </v-card>
<br>
<br>
           <v-text-field
      label="Amount"
      v-model="amount"
      required
    ></v-text-field>
              <v-btn style="float: right" flat>Buy now</v-btn>
    </v-container>
</template>

<script>
export default {
  data() {
    let id = this.$route.params["id"];
    let items = [
      {
        icon: true,
        title: "Every Coin",
        avatar: "/public/1.png"
      },
      { icon: true, title: "F*nny Coin", avatar: "/public/2.png" },
      { title: "Tiki", avatar: "/public/3.png" },
      { title: "Boo Coffee", avatar: "/public/4.png" },
      { title: "Next Movement", avatar: "/public/5.png" },
      { title: "Leflair", avatar: "/public/6.png" }
    ];
    return {
      item: items[id],
      amount: 0,
    };
  },

  mounted() {}
};
</script>
