shader_type canvas_item;

uniform vec4 color = vec4(1, 1, 1, 1);

uniform float strength : hint_range(0.0, 10.0) = 0.0;
uniform vec2 center = vec2(0.5, 0.5);

uniform float piexelFactor : hint_range(0, 10) = 1;

void vertex() {
}

void fragment() {
	vec2 uv;
	// PIXELATION
	vec2 pixelNumber = vec2(textureSize(TEXTURE, 0)) / piexelFactor;
	uv = round(UV * pixelNumber) / pixelNumber;

	// VORTEX
    vec2 delta = uv - center;
    float dist = length(delta);
    float angle = strength * dist;
    float cos_angle = cos(angle);
    float sin_angle = sin(angle);
    vec2 rotated_uv;
    rotated_uv.x = cos_angle * delta.x - sin_angle * delta.y;
    rotated_uv.y = sin_angle * delta.x + cos_angle * delta.y;
    uv = rotated_uv + center;

	// FINAL COLOR
	COLOR = texture(TEXTURE, uv) * color;
}

//void light() {
//	// Called for every pixel for every light affecting the CanvasItem.
//	// Uncomment to replace the default light processing function with this one.
//}
