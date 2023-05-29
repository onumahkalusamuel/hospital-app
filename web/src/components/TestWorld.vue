<script lang="ts">
export default {
    name: 'TestWorld',
    props: {
      msg: String,
    },
    methods: {
      onClick: function () {
        console.log('clicked!');
      },
    },
  };</script>
<template>
  <fluent-card>
    <h2>{{msg}}</h2>
    <fluent-button appearance="accent" v-on:click="onClick">Click Me</fluent-button>
  </fluent-card>
</template>