<template>
  <Toast />

  <div class="col-md-12">
    <div class="card card-container">

      <div class="flex flex-row justify-content-center">
        <img
            id="profile-img"
            src="//ssl.gstatic.com/accounts/ui/avatar_2x.png"
            class="profile-img-card"
        />
      </div>

      <Form @submit="handleLogin" :validation-schema="schema">
        <div class="form-group">
          <label for="username">Имя пользователя</label>
          <Field name="username" type="text" class="form-control" />
          <ErrorMessage name="username" class="error-feedback" />
        </div>

        <div class="form-group">
          <label for="password">Пароль</label>
          <Field name="password" type="password" class="form-control" />
          <ErrorMessage name="password" class="error-feedback" />
        </div>

        <div class="form-group flex flex-row justify-content-center p-2">
          <button class="btn btn-primary btn-block" style="background-color: #04AA6D;" :disabled="loading">
            <span
              v-show="loading"
              class="spinner-border spinner-border-sm"
            ></span>
            <span>Войти</span>
          </button>
        </div>

<!--        <div class="form-group">-->
<!--          <div v-if="message" class="alert alert-danger" role="alert">-->
<!--            {{ message }}-->
<!--          </div>-->
<!--        </div>-->
      </Form>
    </div>
  </div>
</template>

<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";
import { useToast } from "primevue/usetoast";
import Toast from 'primevue/toast';

export default {
  name: "Login",
  components: {
    Form,
    Field,
    ErrorMessage,
    Toast
  },
  data() {
    const schema = yup.object().shape({
      username: yup.string().required("Необходимо ввести имя пользователя"),
      password: yup.string().required("Необходимо ввести пароль"),
    });

    return {
      loading: false,
      message: "",
      schema,
      toast: useToast(),
    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    },
  },
  created() {
    if (this.loggedIn) {
      this.$router.push("/profile");
    }
  },
  methods: {
    handleLogin(user) {
      this.loading = true;

      this.$store.dispatch("auth/login", user).then(
        () => {
          this.$router.push("/salads")
            .then(() => { this.$router.go(0) })
        },
        (error) => {
          this.loading = false;
          this.message =
            (error.response &&
              error.response.data &&
              error.response.data.message) ||
            error.message ||
            error.toString();

          this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `${this.message}`, life: 3000 });
        }
      );
    },
  },
};
</script>

<style scoped>
label {
  display: block;
  margin-top: 10px;
}

.card-container.card {
  max-width: 350px !important;
  padding: 40px 40px;
}

.card {
  background-color: #f7f7f7;
  padding: 20px 25px 30px;
  margin: 0 auto 25px;
  margin-top: 50px;
  -moz-border-radius: 2px;
  -webkit-border-radius: 2px;
  border-radius: 2px;
  -moz-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  -webkit-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
}

.profile-img-card {
  width: 96px;
  height: 96px;
  margin: 0 auto 10px;
  display: block;
  -moz-border-radius: 50%;
  -webkit-border-radius: 50%;
  border-radius: 50%;
}

.error-feedback {
  color: red;
}
</style>