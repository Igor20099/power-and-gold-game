embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"cell\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/atlasses/cell.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.1
  }
}
embedded_components {
  id: "label"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "color {\n"
  "  x: 0.6\n"
  "  y: 0.6\n"
  "  z: 0.0\n"
  "}\n"
  "text: \"\\320\\241\\320\\242\\320\\220\\320\\240\\320\\242\\n"
  "\"\n"
  "  \"\\n"
  "\"\n"
  "  \"\"\n"
  "font: \"/main/assets/fonts/pixel_art.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    z: 0.2
  }
}
