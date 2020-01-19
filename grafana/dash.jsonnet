local dash = (import "grafonnet/grafana.libsonnet") +
{
  dash: $.dashboard.new(
    title="Ferraris KSS",
    description="DESC",
    editable=false,
    graphTooltip=2,
    refresh="5s",
    schemaVersion=20,
    style="dark",
    tags=["asdf", "abc", "3"],
    timezone="utc",
    uid="JrN9MqhZz",
    time_from="now-24h",
    time_to="now",
  )
  // annotations
  .addAnnotation($.annotation.default)
  .addAnnotation($.annotation.datasource(
    name="Test",
    datasource="-- Grafana --"
  ))
  // panels
  .addPanel(
    panel=$.graphPanel.new(
      title="Stromverbrauch",
      datasource="kss",
      fill=5,
    ),
    gridPos={h: 12, w: 23, x: 0, y: 0}
  )
};

dash.dash
