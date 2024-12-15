<template>
  <div>
    <h1>Messages from Golang Backend</h1>
    <ul>
      <!-- ใช้ v-for แสดงข้อความจาก Backend -->
      <li v-for="message in messages" :key="message.id">{{ message.text }}</li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      messages: [] // เก็บข้อมูลที่ดึงมาจาก Backend
    };
  },
  async mounted() {
    try {
      // ดึงข้อมูลจาก Backend Go API
      const response = await fetch("http://localhost:8080/api/messages");
      const data = await response.json();
      console.log("Data from API:", data); // ตรวจสอบข้อมูลใน Console
      this.messages = data; // นำข้อมูลที่ได้มาจาก API ไปแสดงผล
    } catch (error) {
      console.error("Error fetching data:", error); // แสดง Error ในกรณีเกิดปัญหา
    }
  }
};
</script>

<style scoped>
h1 {
  color: #42b983;
  text-align: center;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  font-size: 18px;
  text-align: center;
  margin: 5px 0;
}
</style>
