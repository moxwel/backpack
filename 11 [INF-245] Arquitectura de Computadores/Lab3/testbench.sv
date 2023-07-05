// testbench.sv
module testbench;

  // Valores a usar
  logic [5:0]  addr;
  logic [18:0] instr;
  logic [2:0]  op;
  logic [7:0]  a;
  logic [7:0]  b;
  logic signed [7:0] result; // El resultado esta en complemento 2

  // Modulos y conexiones
  ROM rom(.addr(addr),
          .data(instr));

  SPLITTER splitter(.in(instr),
                    .op(op),
                    .a(a),
                    .b(b));

  ALU alu(.op(op),
          .a(a),
          .b(b),
          .result(result));

  // Codigo
  initial begin
    $display("=======COMIENZO DE CIRCUITO=======");
    $display("Addr   Instr     Op      Num A        Num B        Res");

    // Contador que accede y recorre la ROM
    for (int i = 0; i < 64; i++) begin
      addr = i; #100;

      // Detener for loop cuando se alcanza memoria no definida
      if (instr === 'x) begin
        $display("[ROM][!] Instruccion no definida en: %2h (%d)", addr, addr);
        // $display("----------");
        break;
      end

      // $display("[ROM] Direccion: %b (%d)", addr, addr);
      // $display("[ROM] Instruccion: %h - %b", instr, instr);
      // $display("[SPL] Operacion: %b", op);
      // $display("[SPL] Numero A: %b (%d)", a, a);
      // $display("[SPL] Numero B: %b (%d)", b, b);
      // $display("[ALU] Resultado: %b (%d)", result, result);
      $display("%2h     %h     %b     %b     %b     %b", i, instr, op, a, b, result);
      // $display("----------");
    end

    $display("=========FIN DE CIRCUITO==========");
  end

endmodule
