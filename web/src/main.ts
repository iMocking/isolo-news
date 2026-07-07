import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'
import i18n from './i18n'
import './styles/main.css'
import './assets/article-typography.css'
import { useThemeStore } from './stores'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.use(i18n)

// Initialize theme
const themeStore = useThemeStore()
themeStore.initializeTheme()

app.mount('#app')