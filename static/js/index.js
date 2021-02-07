currentPage = 1
pageSize = 40

idx = getParameters()["page"]
if (idx != null) {
  currentPage = Number(idx)
}

search = "woman"
search = "?search=" + search
fetch('/api/comics/' + String(currentPage) + search)
  .then(response => response.json())
  .then(data => {
    v.$refs.list.comics = data.data
    pageCount = Math.round(data.count / 40 + 0.5)
    // if (pageCount == 0) {
    // pageCount = 1
    // }
    v.$refs.index.pageCount = pageCount
  })
  .catch(error => console.log(error));

Vue.component('comic-list', {
  template: `
  <ul class="comic-list">
    <li class="comic" v-for="info in comics" class="comic" :style="{backgroundImage:'url(/image/'+info.cover+')'}">
      <div class="overlay" :class="{'download': info.download, 'bigsize': info.size > 100}" @click="window.location.href = '/page/comic.html?id='+info.id">
        <h3 @click.stop="">{{ info.name }}</h3>
        <ul>
          <li class="size" >{{ info.size }} M</li>
          <li>{{ info.size }} Mb</li>
          <li>{{ info.pages }} P</li>
          <li>{{ info.year }}</li>
          <li>{{ info.publisher }}</li>
        </ul>
        <div class="download" @click="toggleDownload(info)" @click.stop=""><a :href = "'https://florenfile.com/' + info.download_url" title = "download" target="_blank">Download</a></div>
        <div class="favorite" :class="{'enable': info.favorite}" @click="toggleFavorite(info)" @click.stop="">♥</div>
        <div class="downloaded" :class="{'enable': info.download}" @click="toggleDownload(info)" @click.stop="">⇊</div>
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
    go: (targetUrl => window.location.href = targetUrl),
    toggleFavorite: function (info) {
      op = ''
      url = '/api/comic/favorite/' + String(info.id)
      if (info.favorite) {
        op = 'DELETE'
      } else {
        op = 'POST'
      }
      fetch(url, {
          method: op
        })
        .then(response => response.json())
        .then(data => {
          info.favorite = !info.favorite
        })
        .catch(error => console.log(error));
    },
    toggleDownload: function (info) {
      op = ''
      url = '/api/comic/download/' + String(info.id)
      if (info.download) {
        op = 'DELETE'
      } else {
        op = 'POST'
      }
      fetch(url, {
          method: op
        })
        .then(response => response.json())
        .then(data => {
          info.download = !info.download
        })
        .catch(error => console.log(error));
    }
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


Vue.component('gopage', {
  template: `
  <div class="go">
    <input class="go-input" type="text" id="fname" name="page">
    <label for="page">/60</label>
  </div>
  `,
  data: function () {},
  methods: {

  }
})

function getParameters() {
  var vars = {};
  var parts = window.location.href.replace(/[?&]+([^=&]+)=([^&]*)/gi, function (m, key, value) {
    vars[key] = value;
  });
  return vars;
}

function handleKeydown(e) {
  switch (e.keyCode) {
    case 37: // <-
      if (currentPage == 1) {
        return
      }
      window.location.href = '/page/index.html' + '?page=' + String(currentPage - 1)
      break
    case 39: // ->
      if (currentPage == pageCount) {
        return
      }
      window.location.href = '/page/index.html' + '?page=' + String(currentPage + 1)
      break;
  }
}
