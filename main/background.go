components {
  id: "background"
  component: "/scripts/background.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Purple_Nebula_04-1024x1024\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/background.atlas\"\n"
  "}\n"
  ""
  position {
    x: 512.0
    y: 512.0
  }
}
