import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

// import router functions
import { createRouter, createWebHistory } from 'vue-router'
// setup routes
const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'Recipes',
      component: () => import('./components/RecipeList/RecipeList.vue')
    },
    {
      path: '/add-recipe',
      name: 'Add Recipe',
      component: () => import('./components/AddRecipe/AddRecipe.vue')
    },
    {
      path: '/saved-recipes',
      name: 'Saved Recipes',
      component: () => import('./components/SavedRecipes/SavedRecipes.vue')
    },
  ]
})

const app = createApp(App)

// tell Vue to use router
app.use(router)
app.mount('#app')