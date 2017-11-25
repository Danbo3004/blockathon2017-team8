<template>
    <v-app id="example-1">
        <v-navigation-drawer persistent v-model="drawer">
            <v-list>
                <v-list-tile>
                    <v-list-tile-content>
                        <v-list-tile-title >
                            <span>Menu</span>
                        </v-list-tile-title>
                    </v-list-tile-content>
                </v-list-tile>
                <v-divider></v-divider>
                <template v-for="(item, index) in items">
                    <v-list-tile :href="item.href" :to="{name: item.href}">
                        <v-list-tile-action>
                            <v-icon light v-html="item.icon"></v-icon>
                        </v-list-tile-action>
                        <v-list-tile-content>
                            <v-list-tile-title v-html="item.title"></v-list-tile-title>
                        </v-list-tile-content>
                    </v-list-tile>
                </template>
            </v-list>
        </v-navigation-drawer>
        <v-toolbar class="amber darken-3" dark>
            <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
            <v-toolbar-title>Blocker App</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn :to="{name: 'home'}">
                <v-icon>home</v-icon>
            </v-btn>
        </v-toolbar>
        <main>
            <v-fade-transition mode="out-in">
                <router-view></router-view>
            </v-fade-transition>
        </main>
        <!--
        <v-footer class="amber darken-3" dark>
            <span>© 2017 - Blocker</span>
        </v-footer>
        -->
    </v-app>
</template>

<script type="text/babel">
  export default {
    data() {
      return {
        drawer: true,
        items: [{
          href: 'home',
          router: true,
          title: 'Home',
          icon: 'home',
        }, {
          href: 'register',
          router: true,
          title: 'Register',
          icon: 'chrome_reader_mode',
        }, {
          href: 'list_company',
          router: true,
          title: 'Buy Tokens',
          icon: 'shopping_basket',
        }, {
          href: 'transfer',
          router: true,
          title: 'Transfer Tokens',
          icon: 'payment',
        }],
      };
    },
    methods: {
      openGithub() {

      },
    },
  };
</script>

<style lang="stylus">
    @import '../node_modules/vuetify/src/stylus/main';
    @import 'css/main.css';
</style>
