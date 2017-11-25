<template>
  <v-tabs fixed centered>
    <v-toolbar extended color="cyan">
      <v-toolbar-title slot="extension" class="display-2">Register<br>&nbsp;</v-toolbar-title>
    </v-toolbar>
    <v-tabs-bar style="background: #d6d6d6" dark>
      <v-tabs-slider color="yellow"></v-tabs-slider>
      <v-tabs-item :href="'#tab-r1'">
        Customer
      </v-tabs-item>
      <v-tabs-item :href="'#tab-r2'">
        Company
      </v-tabs-item>
    </v-tabs-bar>
    <v-tabs-items>
      <v-tabs-content :id="'tab-r1'">
        <v-card flat>
          <v-card-text>
            <v-form v-model="valid">
                <v-text-field
                label="Name"
                v-model="name"
                :rules="nameRules"
                :counter="10"
                required
                ></v-text-field>
                <v-text-field
                label="E-mail"
                v-model="email"
                :rules="emailRules"
                required
                ></v-text-field>
                <v-text-field
                label="Password"
                v-model="password"
                required
                ></v-text-field>
            </v-form>
          </v-card-text>
        </v-card>
      </v-tabs-content>
      <v-tabs-content :id="'tab-r2'">
        <v-card flat>
          <v-card-text>
            <v-form v-model="valid">
                <v-text-field
                label="Name"
                v-model="name"
                :rules="nameRules"
                :counter="10"
                required
                ></v-text-field>
                <v-text-field
                label="E-mail"
                v-model="email"
                :rules="emailRules"
                required
                ></v-text-field>
                <v-text-field
                label="Password"
                v-model="password"
                required
                ></v-text-field>
                 <v-text-field
                label="Token Name"
                v-model="tokenName"
                required
                ></v-text-field>
                <v-text-field
                label="Amount"
                v-model="tokenAmount"
                required
                ></v-text-field>
            </v-form>
          </v-card-text>
        </v-card>
      </v-tabs-content>
    </v-tabs-items>
    <v-card flat>
          <v-card-text>
            <div>
              <v-btn right color="amber" style="color: #fff; background: #ff6f00; float: right">Register</v-btn>
              <div style="clear:both"></div>
            </div>
          </v-card-text>
        </v-card>
  </v-tabs>
</template>

<script>
export default {
  data() {
    return {
        valid: false,
        name: '',
        nameRules: [
          (v) => !!v || 'Name is required',
          (v) => v.length <= 10 || 'Name must be less than 10 characters'
        ],
        email: '',
        emailRules: [
          (v) => !!v || 'E-mail is required',
          (v) => /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid'
        ],
        tokenName: '',
        tokenAmount: 0,
        password: ''
      }
  },

  mounted() {}
};
</script>
