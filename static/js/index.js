currentPage = 1
pageSize = 40

idx = getParameters()["page"]
if (idx != null) {
  currentPage = Number(idx)
}

fetch('/api/comics/' + String(currentPage))
  .then(response => response.json())
  .then(data => {
    v.$refs.list.comics = data.data
    v.$refs.index.pageCount = Math.round(data.count / 40)
  })
  .catch(error => console.log(error));

Vue.component('comic-list', {
  template: `
  <ul class="comic-list">
    <li v-for="info in comics" class="comic" :style="{backgroundImage:'url(/image/'+info.cover+')'}">
      <div class="overlay">
        {{ info.name }}
        <p>Size: {{ info.size }} Mb
        <p>Pages: {{ info.pages }}
        <p>Year: {{ info.year }}</p>
        <a :href = "'https://florenfile.com/' + info.download_url" title = "download" > Download </a>
      </div>
    </li>
  </ul>
  `,
  data: function () {
    return {
      comics: [],
    }
  },
  methods: {
    go: (targetUrl => window.location.href = targetUrl)
  }
})

Vue.component('paginate', VuejsPaginate)

Vue.component('index', {
  template: `
    <paginate v-model="page" :pageCount="pageCount" :page-range="10" :margin-pages="1" :containerClass="'pagination'" :clickHandler="clickCallback"></paginate>
  `,
  data: function () {
    return {
      pageCount: 0,
      page: currentPage
    }
  },
  created: function () {
    // v.$refs.index.page = 3
  },
  methods: {
    clickCallback: function (p) {
      window.location.href = '/page/index.html' + '?page=' + p
    }
  }
})

function getParameters() {
  var vars = {};
  var parts = window.location.href.replace(/[?&]+([^=&]+)=([^&]*)/gi, function (m, key, value) {
    vars[key] = value;
  });
  return vars;
}
