#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <ranges>
#include <vector>
#include <algorithm>
#include <numeric>

using namespace std::literals::string_view_literals;

int main() {
    std::ifstream input("input.txt");
    std::stringstream buf;
    buf << input.rdbuf();
    auto str = buf.str();

    auto strategy = [&](auto&& fn) {
        auto sums = str
                    | std::views::split("\n"sv)
                    | std::views::transform([&](const auto& s) {
            char o = *(s.begin());
            char m = *(s.begin() + 2);
            return fn(o, m);
        });
        return std::accumulate(sums.begin(), sums.end(), 0);
    };

    std::cout << strategy([](char o, char m) {
        bool win = (o == 'A' && m == 'Y') || (o == 'B' && m == 'Z') || (o == 'C' && m == 'X');
        return ((o - 'A') == (m - 'X') ? 3 : win ? 6 : 0) + (m == 'X' ? 1 : m == 'Y' ? 2 : 3);
    }) << '\n';

    std::cout << strategy([](char o, char m) {
        if (m == 'X') {
            if (o == 'A') return 3;
            if (o == 'B') return 1;
            if (o == 'C') return 2;
        } else if (m == 'Y') {
            return 3 + (o - 'A' + 1);
        } else { // Z
            if (o == 'A') return 6 + 2;
            if (o == 'B') return 6 + 3;
            if (o == 'C') return 6 + 1;
        }
    });

    return 0;
}
