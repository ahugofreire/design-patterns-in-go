package main

import "fmt"

// Interface do cliente
type MediaPlayer interface {
	play(audioType string, fileName string)
}

// Interface do adaptador
type AdvancedMediaPlayer interface {
	playVlc(fileName string)
	playMp4(fileName string)
}

// Implementação do adaptador Vlc
type VlcPlayer struct{}

func (v *VlcPlayer) playVlc(fileName string) {
	fmt.Printf("Playing vlc file. Name: %s\n", fileName)
}

func (v *VlcPlayer) playMp4(fileName string) {
	// Não faz nada
}

// Implementação do adaptador Mp4
type Mp4Player struct{}

func (m *Mp4Player) playMp4(fileName string) {
	fmt.Printf("Playing mp4 file. Name: %s\n", fileName)
}

func (m *Mp4Player) playVlc(fileName string) {
	// Não faz nada
}

// Adaptador
type MediaAdapter struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	if audioType == "vlc" {
		return &MediaAdapter{
			advancedMediaPlayer: &VlcPlayer{},
		}
	} else if audioType == "mp4" {
		return &MediaAdapter{
			advancedMediaPlayer: &Mp4Player{},
		}
	} else {
		return nil
	}
}

func (m *MediaAdapter) play(audioType string, fileName string) {
	if audioType == "vlc" {
		m.advancedMediaPlayer.playVlc(fileName)
	} else if audioType == "mp4" {
		m.advancedMediaPlayer.playMp4(fileName)
	}
}

// Implementação do MediaPlayer utilizando o adaptador
type AudioPlayer struct {
	mediaAdapter *MediaAdapter
}

func NewAudioPlayer() *AudioPlayer {
	return &AudioPlayer{}
}

func (a *AudioPlayer) play(audioType string, fileName string) {
	if audioType == "mp3" {
		fmt.Printf("Playing mp3 file. Name: %s\n", fileName)
	} else if audioType == "vlc" || audioType == "mp4" {
		a.mediaAdapter = NewMediaAdapter(audioType)
		a.mediaAdapter.play(audioType, fileName)
	} else {
		fmt.Printf("Invalid media. %s format not supported\n", audioType)
	}
}

func main() {
	audioPlayer := NewAudioPlayer()

	audioPlayer.play("mp3", "beyond_the_horizon.mp3")
	audioPlayer.play("mp4", "alone.mp4")
	audioPlayer.play("vlc", "far_far_away.vlc")
	audioPlayer.play("avi", "mind_me.avi")
}
