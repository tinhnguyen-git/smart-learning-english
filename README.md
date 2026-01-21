# smart-learning-english



Hướng dẫn sử dụng sau khi chạy Script
Sau khi script chạy xong, dự án của bạn đã có đủ 4 trụ cột 10/10:

Bộ nhớ (Memory): File project_brain.json đã được tạo. Từ giờ, hãy bắt đầu mọi prompt bằng: "Đọc project_brain.json và cho tôi biết tình trạng dự án..."

An toàn (Sandboxing): Thư mục .devcontainer đã sẵn sàng. Khi mở Antigravity, nó sẽ hỏi bạn có muốn mở trong Container không. Hãy chọn "Reopen in Container". Lúc này, Agent sẽ bị nhốt trong hộp Docker, không thể xóa nhầm file hệ thống của bạn.

Di động (Mobile): Script setup_android.sh sẽ tự động cài Android Command Line Tools bên trong Container, cho phép Agent chạy lệnh flutter build apk mà không cần bạn cài Android Studio nặng nề trên máy thật.

Hạ tầng (Infra): File infra/docker/docker-compose.yml đã sẵn sàng để Agent bật Database PostgreSQL lên chỉ với 1 lệnh.

## Mobile Development Setup
**Prerequisites:**
- Flutter SDK is installed locally in `workspace/smart-learning-english/flutter` (since it wasn't available in the environment).
- Add the following to your path or use the direct path to run commands:
  ```bash
  export PATH="$PATH:$(pwd)/flutter/bin"
  ```

## Project Structure
- `backend/`: Go
- `mobile/`: Flutter (Clean Architecture)