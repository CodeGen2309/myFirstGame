components {
  id: "ship"
  component: "/scripts/ship.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"New Animation\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/ships.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "bullet"
  type: "factory"
  data: "prototype: \"/main/missile.go\"\n"
  ""
}
