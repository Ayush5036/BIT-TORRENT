# BitTorrent

## Introduction

Welcome to the BitTorrent project! This project aims to implement a basic BitTorrent client in GoLang, allowing users to share and download files using the BitTorrent protocol. Before you start, please make sure you have GoLang set up on your device. If not, follow the installation instructions for GoLang at https://golang.org/ to get started.

## Cloning the Repository

To clone this GitHub repository, use the following command:

```
git clone https://github.com/Suggest-Name/BIT-TORRENT.git
```

## Running the BitTorrent Client

To run the BitTorrent client after installation, use the following command:

```
go run . "path to torrent" "path to download destination" > output.txt 2> error.txt
```

For example:

```
go run . torrent/0.torrent download > output.txt 2> error.txt
```

Replace "torrent/0.torrent" with the path to the .torrent file you want to download, and "download" with the folder where you want to save the downloaded files.

This command will execute the BitTorrent client, and the standard output will be redirected to the "output.txt" file, while the standard error will be redirected to the "error.txt" file for better logging.

## Currently Working On

- NAT Traversal
- Distributed Hash Table (DHT)
- Supporting Magnet Links using external packages

## Contact

If you have any questions, suggestions, or need assistance, you can reach out to us through the GitHub repository's issue section.

Thank you for your interest in the BitTorrent project! Happy sharing and downloading!
