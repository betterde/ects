<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-body" :class="classes">
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Dashboard",
    data() {
      return {
        classes: ['animated', 'fade-in', 'fast']
      }
    }
  }
</script>

<style lang="scss">

</style>
