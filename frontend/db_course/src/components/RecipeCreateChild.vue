<template>
  <div class="card flex flex-column md:flex-row gap-3">
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-pencil"></i>
      </InputGroupAddon>
      <InputNumber v-model="nos" placeholder="Количество порций" :class="{ 'p-invalid': nos <= 0 }" />
      <InputNumber v-model="ttc" placeholder="Время приготовления (мин.)" :class="{ 'p-invalid': ttc <= 0 }" />
      <Button @click="createRecipe" style="border-top-right-radius: 4px; border-bottom-right-radius: 4px; width: 108px">
         Создать
      </Button>
    </InputGroup>
  </div>
</template>

<script>
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputNumber from 'primevue/inputnumber';
import Button from 'primevue/button';
import { ref } from 'vue';
import { useToast } from "primevue/usetoast";
import Toast from 'primevue/toast';

import RecipeModel from "@/models/RecipeModel.js";
import RecipeService from "@/services/recipe.service.js";

export default {
  name: 'CreateRecipeChild',

  props: {
    salad_id: String,
  },

  components: {
    Button,
    InputGroup,
    InputGroupAddon,
    InputNumber,
    Toast,
  },
  setup() {
    const nos = ref(5)
    const ttc = ref(60)

    return {
      nos,
      ttc,
      toast: useToast(),
    };
  },
  created() {
  },
  methods: {
    createRecipe() {
      const recipe = new RecipeModel(this.nos, this.ttc, this.salad_id)

      RecipeService.createRecipe(recipe)
          .then(response => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Рецепт успешно создан', life: 3000 });
            this.creatingSuccess(response.data.data.recipe)
          })
          .catch(error => {
            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при создании рецепта: ${error.response.data.error}`, life: 3000 });
            this.creatingError(error.response.data.error)
          })
    },

    creatingSuccess(recipe) {
      this.$emit('creatingSuccess', recipe)
    },

    creatingError(errorText) {
      this.$emit('creatingError', errorText)
    }
  }
}

</script>