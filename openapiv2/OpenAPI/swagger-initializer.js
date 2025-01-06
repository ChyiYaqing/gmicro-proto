window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    urls: [
      {url: "user/v1/user.swagger.json", name: "User API v1"},
      {url: "order/v1/order.swagger.json", name: "Order API v1"},
      {url: "payment/v1/payment.swagger.json", name: "Payment API v1"},
      {url: "shipping/v1/shipping.swagger.json", name: "Shipping API v1"},
    ],
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });

  //</editor-fold>
};
