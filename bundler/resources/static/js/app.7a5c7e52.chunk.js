(this.webpackJsonp=this.webpackJsonp||[]).push([[0],{112:function(t,e){},70:function(t,e,n){"use strict";n.r(e),n.d(e,"default",(function(){return C}));var r=n(37),a=n.n(r),i=n(19),c=n.n(i),l=n(1),u=n(95),o=n(27),d=n(88),s=n(12),f=n(7),h=n(96),p=n(94),b=n(71),j=n.n(b),x=n(57),y=n(97),v=n(72),g=n.n(v),m=n(11),k=function(){var t=c()((function*(){try{var t=j.a.stringify({grant_type:"client_credentials"}),e=x.Buffer.from("e76ef41e470a45cb923f6f77aaae5bf1:f170d1874c8d46d3911f6066b6ac822a").toString("base64");return(yield y.default.post("https://accounts.spotify.com/api/token",t,{headers:{Authorization:"Basic "+e,"Content-Type":"application/x-www-form-urlencoded"}})).data.access_token}catch(n){console.log(n)}}));return function(){return t.apply(this,arguments)}}(),O=function(){var t=c()((function*(t,e){var n="https://api.spotify.com/v1/search?q="+e+"&type=track&market=US&limit=10";try{return(yield y.default.get(n,{headers:{Authorization:"Bearer "+t}})).data}catch(r){console.log(r)}}));return function(e,n){return t.apply(this,arguments)}}();function C(){var t=l.useState(""),e=a()(t,2),n=e[0],r=e[1],i=l.useState(""),f=a()(i,2),p=f[0],b=f[1],j=l.useState([]),x=a()(j,2),y=x[0],v=x[1],C=l.useCallback((function(){k().then((function(t){b(t)}))}),[]);l.useEffect((function(){C();var t=setInterval((function(){C()}),36e5);return function(){clearInterval(t)}}),[C]);var T=l.useCallback(function(){var t=c()((function*(t){var e=yield O(p,t);if(void 0===e)return v([]);var n=e.tracks.items;v(n.map((function(t){return{id:t.id,title:t.name,artist:t.artists[0].name,album:t.album.name,link:t.uri}})))}));return function(e){return t.apply(this,arguments)}}(),[p]),w=g()(T,500),B=l.useCallback(function(){var t=c()((function*(t){w(t),r(t)}));return function(e){return t.apply(this,arguments)}}(),[w]),A=l.useCallback(function(){var t=c()((function*(t){console.log(t),v([]),r("")}));return function(e){return t.apply(this,arguments)}}(),[]);return Object(m.jsxs)(s.default,{style:S.container,children:[Object(m.jsxs)(s.default,{children:[Object(m.jsx)(o.default,{children:"Token:"}),Object(m.jsx)(o.default,{style:{color:"#666"},children:p}),Object(m.jsx)(s.default,{style:S.card,children:Object(m.jsx)(d.default,{style:S.search,value:n,placeholderTextColor:"#888",placeholder:"Search...",onChangeText:B})})]}),!!y.length&&Object(m.jsx)(s.default,{style:{minHeight:100,maxHeight:400,height:"30%",borderWidth:2,borderTopWidth:0,borderColor:"#ccc"},children:Object(m.jsx)(u.default,{children:y.map((function(t){return Object(m.jsxs)(m.Fragment,{children:[Object(m.jsx)(s.default,{style:S.divider}),Object(m.jsx)(h.default,{onPress:function(){return A(t.link)},children:Object(m.jsxs)(s.default,{style:S.card,pointerEvents:"auto",children:[Object(m.jsxs)(o.default,{children:["#",t.id]}),Object(m.jsxs)(o.default,{children:[t.artist," - ",t.title]}),Object(m.jsx)(o.default,{children:"Album: "+t.album}),Object(m.jsx)(o.default,{children:"Link: "+t.link})]})})]})}))})})]})}var S=f.default.create({container:{flex:1,display:"flex",paddingTop:p.default.statusBarHeight,backgroundColor:"#ecf0f1",padding:8},search:{borderWidth:1,borderRadius:4,padding:4,flex:1},divider:{paddingVertical:4},card:{padding:16,backgroundColor:"#fff",borderRadius:4}})},98:function(t,e,n){t.exports=n(146)}},[[98,1,2]]]);
//# sourceMappingURL=app.7a5c7e52.chunk.js.map