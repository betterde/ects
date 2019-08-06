<template>
<h1>Hello</h1>
</template>

<script>
  export default {
    name: "PipelineDetail",
    data() {
      return {

      }
    },
    mounted() {
      window.console.log(this.$route.params.id)
    }
  }
</script>

<style lang="scss">

</style>
