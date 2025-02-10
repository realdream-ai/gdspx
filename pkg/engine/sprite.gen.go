//go:build !pure_engine
// +build !pure_engine

/*------------------------------------------------------------------------------
//   This code was generated by template sprite.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "sprite.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

package engine

import (
	. "github.com/realdream-ai/mathf"
)

func (pself *Sprite) AddForce(force Vec2) {
	SpriteMgr.AddForce(pself.Id, force)
}

func (pself *Sprite) AddImpulse(impulse Vec2) {
	SpriteMgr.AddImpulse(pself.Id, impulse)
}

func (pself *Sprite) CheckCollision(target Object, is_src_trigger bool, is_dst_trigger bool) bool {
	return SpriteMgr.CheckCollision(pself.Id, target, is_src_trigger, is_dst_trigger)
}

func (pself *Sprite) CheckCollisionWithPoint(point Vec2, is_trigger bool) bool {
	return SpriteMgr.CheckCollisionWithPoint(pself.Id, point, is_trigger)
}

func (pself *Sprite) CloneSprite() Object {
	return SpriteMgr.CloneSprite(pself.Id)
}

func (pself *Sprite) CreateSprite(path string) Object {
	return SpriteMgr.CreateSprite(path)
}

func (pself *Sprite) DestroySprite() bool {
	return SpriteMgr.DestroySprite(pself.Id)
}

func (pself *Sprite) GetAnim() string {
	return SpriteMgr.GetAnim(pself.Id)
}

func (pself *Sprite) GetAnimFrame() int64 {
	return SpriteMgr.GetAnimFrame(pself.Id)
}

func (pself *Sprite) GetAnimOffset() Vec2 {
	return SpriteMgr.GetAnimOffset(pself.Id)
}

func (pself *Sprite) GetAnimPlayingSpeed() float64 {
	return SpriteMgr.GetAnimPlayingSpeed(pself.Id)
}

func (pself *Sprite) GetAnimSpeedScale() float64 {
	return SpriteMgr.GetAnimSpeedScale(pself.Id)
}

func (pself *Sprite) GetChildPosition(path string) Vec2 {
	return SpriteMgr.GetChildPosition(pself.Id, path)
}

func (pself *Sprite) GetChildRotation(path string) float64 {
	return SpriteMgr.GetChildRotation(pself.Id, path)
}

func (pself *Sprite) GetChildScale(path string) Vec2 {
	return SpriteMgr.GetChildScale(pself.Id, path)
}

func (pself *Sprite) GetCollisionLayer() int64 {
	return SpriteMgr.GetCollisionLayer(pself.Id)
}

func (pself *Sprite) GetCollisionMask() int64 {
	return SpriteMgr.GetCollisionMask(pself.Id)
}

func (pself *Sprite) GetColor() Color {
	return SpriteMgr.GetColor(pself.Id)
}

func (pself *Sprite) GetFloorNormal() Vec2 {
	return SpriteMgr.GetFloorNormal(pself.Id)
}

func (pself *Sprite) GetGravity() float64 {
	return SpriteMgr.GetGravity(pself.Id)
}

func (pself *Sprite) GetLastMotion() Vec2 {
	return SpriteMgr.GetLastMotion(pself.Id)
}

func (pself *Sprite) GetMass() float64 {
	return SpriteMgr.GetMass(pself.Id)
}

func (pself *Sprite) GetPosition() Vec2 {
	return SpriteMgr.GetPosition(pself.Id)
}

func (pself *Sprite) GetPositionDelta() Vec2 {
	return SpriteMgr.GetPositionDelta(pself.Id)
}

func (pself *Sprite) GetRealVelocity() Vec2 {
	return SpriteMgr.GetRealVelocity(pself.Id)
}

func (pself *Sprite) GetRenderScale() Vec2 {
	return SpriteMgr.GetRenderScale(pself.Id)
}

func (pself *Sprite) GetRotation() float64 {
	return SpriteMgr.GetRotation(pself.Id)
}

func (pself *Sprite) GetScale() Vec2 {
	return SpriteMgr.GetScale(pself.Id)
}

func (pself *Sprite) GetTexture() string {
	return SpriteMgr.GetTexture(pself.Id)
}

func (pself *Sprite) GetTriggerLayer() int64 {
	return SpriteMgr.GetTriggerLayer(pself.Id)
}

func (pself *Sprite) GetTriggerMask() int64 {
	return SpriteMgr.GetTriggerMask(pself.Id)
}

func (pself *Sprite) GetVelocity() Vec2 {
	return SpriteMgr.GetVelocity(pself.Id)
}

func (pself *Sprite) GetVisible() bool {
	return SpriteMgr.GetVisible(pself.Id)
}

func (pself *Sprite) GetWallNormal() Vec2 {
	return SpriteMgr.GetWallNormal(pself.Id)
}

func (pself *Sprite) GetZIndex() int64 {
	return SpriteMgr.GetZIndex(pself.Id)
}

func (pself *Sprite) IsAnimCentered() bool {
	return SpriteMgr.IsAnimCentered(pself.Id)
}

func (pself *Sprite) IsAnimFlippedH() bool {
	return SpriteMgr.IsAnimFlippedH(pself.Id)
}

func (pself *Sprite) IsAnimFlippedV() bool {
	return SpriteMgr.IsAnimFlippedV(pself.Id)
}

func (pself *Sprite) IsCollisionEnabled() bool {
	return SpriteMgr.IsCollisionEnabled(pself.Id)
}

func (pself *Sprite) IsOnCeiling() bool {
	return SpriteMgr.IsOnCeiling(pself.Id)
}

func (pself *Sprite) IsOnCeilingOnly() bool {
	return SpriteMgr.IsOnCeilingOnly(pself.Id)
}

func (pself *Sprite) IsOnFloor() bool {
	return SpriteMgr.IsOnFloor(pself.Id)
}

func (pself *Sprite) IsOnFloorOnly() bool {
	return SpriteMgr.IsOnFloorOnly(pself.Id)
}

func (pself *Sprite) IsOnWall() bool {
	return SpriteMgr.IsOnWall(pself.Id)
}

func (pself *Sprite) IsOnWallOnly() bool {
	return SpriteMgr.IsOnWallOnly(pself.Id)
}

func (pself *Sprite) IsPlayingAnim() bool {
	return SpriteMgr.IsPlayingAnim(pself.Id)
}

func (pself *Sprite) IsSpriteAlive() bool {
	return SpriteMgr.IsSpriteAlive(pself.Id)
}

func (pself *Sprite) IsTriggerEnabled() bool {
	return SpriteMgr.IsTriggerEnabled(pself.Id)
}

func (pself *Sprite) MoveAndSlide() {
	SpriteMgr.MoveAndSlide(pself.Id)
}

func (pself *Sprite) PauseAnim() {
	SpriteMgr.PauseAnim(pself.Id)
}

func (pself *Sprite) PlayAnim(p_name string, p_speed float64, isLoop bool, p_revert bool) {
	SpriteMgr.PlayAnim(pself.Id, p_name, p_speed, isLoop, p_revert)
}

func (pself *Sprite) PlayBackwardsAnim(p_name string) {
	SpriteMgr.PlayBackwardsAnim(pself.Id, p_name)
}

func (pself *Sprite) SetAnim(p_name string) {
	SpriteMgr.SetAnim(pself.Id, p_name)
}

func (pself *Sprite) SetAnimCentered(p_center bool) {
	SpriteMgr.SetAnimCentered(pself.Id, p_center)
}

func (pself *Sprite) SetAnimFlipH(p_flip bool) {
	SpriteMgr.SetAnimFlipH(pself.Id, p_flip)
}

func (pself *Sprite) SetAnimFlipV(p_flip bool) {
	SpriteMgr.SetAnimFlipV(pself.Id, p_flip)
}

func (pself *Sprite) SetAnimFrame(p_frame int64) {
	SpriteMgr.SetAnimFrame(pself.Id, p_frame)
}

func (pself *Sprite) SetAnimOffset(p_offset Vec2) {
	SpriteMgr.SetAnimOffset(pself.Id, p_offset)
}

func (pself *Sprite) SetAnimSpeedScale(p_speed_scale float64) {
	SpriteMgr.SetAnimSpeedScale(pself.Id, p_speed_scale)
}

func (pself *Sprite) SetChildPosition(path string, pos Vec2) {
	SpriteMgr.SetChildPosition(pself.Id, path, pos)
}

func (pself *Sprite) SetChildRotation(path string, rot float64) {
	SpriteMgr.SetChildRotation(pself.Id, path, rot)
}

func (pself *Sprite) SetChildScale(path string, scale Vec2) {
	SpriteMgr.SetChildScale(pself.Id, path, scale)
}

func (pself *Sprite) SetColliderCapsule(center Vec2, size Vec2) {
	SpriteMgr.SetColliderCapsule(pself.Id, center, size)
}

func (pself *Sprite) SetColliderCircle(center Vec2, radius float64) {
	SpriteMgr.SetColliderCircle(pself.Id, center, radius)
}

func (pself *Sprite) SetColliderRect(center Vec2, size Vec2) {
	SpriteMgr.SetColliderRect(pself.Id, center, size)
}

func (pself *Sprite) SetCollisionEnabled(enabled bool) {
	SpriteMgr.SetCollisionEnabled(pself.Id, enabled)
}

func (pself *Sprite) SetCollisionLayer(layer int64) {
	SpriteMgr.SetCollisionLayer(pself.Id, layer)
}

func (pself *Sprite) SetCollisionMask(mask int64) {
	SpriteMgr.SetCollisionMask(pself.Id, mask)
}

func (pself *Sprite) SetColor(color Color) {
	SpriteMgr.SetColor(pself.Id, color)
}

func (pself *Sprite) SetDontDestroyOnLoad() {
	SpriteMgr.SetDontDestroyOnLoad(pself.Id)
}

func (pself *Sprite) SetGravity(gravity float64) {
	SpriteMgr.SetGravity(pself.Id, gravity)
}

func (pself *Sprite) SetMass(mass float64) {
	SpriteMgr.SetMass(pself.Id, mass)
}

func (pself *Sprite) SetPhysicProcess(is_on bool) {
	SpriteMgr.SetPhysicProcess(pself.Id, is_on)
}

func (pself *Sprite) SetPosition(pos Vec2) {
	SpriteMgr.SetPosition(pself.Id, pos)
}

func (pself *Sprite) SetProcess(is_on bool) {
	SpriteMgr.SetProcess(pself.Id, is_on)
}

func (pself *Sprite) SetRenderScale(scale Vec2) {
	SpriteMgr.SetRenderScale(pself.Id, scale)
}

func (pself *Sprite) SetRotation(rot float64) {
	SpriteMgr.SetRotation(pself.Id, rot)
}

func (pself *Sprite) SetScale(scale Vec2) {
	SpriteMgr.SetScale(pself.Id, scale)
}

func (pself *Sprite) SetTexture(path string) {
	SpriteMgr.SetTexture(pself.Id, path)
}

func (pself *Sprite) SetTextureAltas(path string, rect2 Rect2) {
	SpriteMgr.SetTextureAltas(pself.Id, path, rect2)
}

func (pself *Sprite) SetTriggerCapsule(center Vec2, size Vec2) {
	SpriteMgr.SetTriggerCapsule(pself.Id, center, size)
}

func (pself *Sprite) SetTriggerCircle(center Vec2, radius float64) {
	SpriteMgr.SetTriggerCircle(pself.Id, center, radius)
}

func (pself *Sprite) SetTriggerEnabled(trigger bool) {
	SpriteMgr.SetTriggerEnabled(pself.Id, trigger)
}

func (pself *Sprite) SetTriggerLayer(layer int64) {
	SpriteMgr.SetTriggerLayer(pself.Id, layer)
}

func (pself *Sprite) SetTriggerMask(mask int64) {
	SpriteMgr.SetTriggerMask(pself.Id, mask)
}

func (pself *Sprite) SetTriggerRect(center Vec2, size Vec2) {
	SpriteMgr.SetTriggerRect(pself.Id, center, size)
}

func (pself *Sprite) SetTypeName(type_name string) {
	SpriteMgr.SetTypeName(pself.Id, type_name)
}

func (pself *Sprite) SetVelocity(velocity Vec2) {
	SpriteMgr.SetVelocity(pself.Id, velocity)
}

func (pself *Sprite) SetVisible(visible bool) {
	SpriteMgr.SetVisible(pself.Id, visible)
}

func (pself *Sprite) SetZIndex(z int64) {
	SpriteMgr.SetZIndex(pself.Id, z)
}

func (pself *Sprite) StopAnim() {
	SpriteMgr.StopAnim(pself.Id)
}
